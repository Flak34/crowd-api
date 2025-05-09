package project_service

import (
	"context"
	"encoding/csv"
	"github.com/Flak34/crowd-api/internal/errors/storage_errors"
	project_model "github.com/Flak34/crowd-api/internal/project/model"
	task_model "github.com/Flak34/crowd-api/internal/task/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"io"
)

func (s *Service) UploadProjectData(ctx context.Context, projectID int, dataReader io.Reader) error {
	db := s.ep.GetDB()
	project, err := s.projectRepo.GetProject(ctx, db, projectID)
	if err != nil {
		if storage_errors.IsNotFound(err) {
			return errors.Wrap(err, "project does not exist")
		}
		return err
	}
	csvReader := csv.NewReader(dataReader)
	inputDataNameToData := make(map[string]project_model.InputData, len(project.Config.InputData))
	for _, data := range project.Config.InputData {
		inputDataNameToData[data.Name] = data
	}
	headerLine, err := csvReader.Read()
	if err != nil {
		return err
	}
	lines, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	outputData := make([]task_model.OutputData, 0, len(project.Config.OutputData))
	for _, data := range project.Config.OutputData {
		outputData = append(outputData, task_model.OutputData{
			Name:            data.Name,
			Type:            string(data.Type),
			Value:           task_model.GetDefaultDataValue(string(data.Type)),
			WithAggregation: data.WithAggregation,
		})
	}
	err = s.ep.TxWrapper(ctx, func(ctx context.Context, tx pgx.Tx) error {
		tasks := make([]task_model.Task, 0, len(lines))
		for _, line := range lines {
			newTask := task_model.Task{
				ProjectID:     project.ID,
				TargetOverlap: project.TargetOverlap,
				OutputData:    outputData,
			}
			inputData := make([]task_model.InputData, 0, len(project.Config.InputData))
			for i, dataValue := range line {
				dataName := headerLine[i]
				data := inputDataNameToData[dataName]
				inputData = append(inputData, task_model.InputData{
					Name:  dataName,
					Type:  string(data.Type),
					Value: dataValue,
				})
			}
			newTask.InputData = inputData
			tasks = append(tasks, newTask)
		}
		err = s.taskRepo.CreateTasks(ctx, tx, tasks...)
		if err != nil {
			return err
		}
		err = s.projectRepo.UpdateStatus(ctx, tx, project.ID, project_model.ProjectActiveStatusName)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
