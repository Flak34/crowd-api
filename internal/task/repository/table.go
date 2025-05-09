package task_repository

import (
	model "github.com/Flak34/crowd-api/internal/task/model"
	"github.com/samber/lo"
	"time"
)

type TaskTable struct {
	ID                  int             `db:"id"`
	ProjectID           int             `db:"project_id"`
	TargetOverlap       int             `db:"target_overlap"`
	CurrentOverlap      int             `db:"current_overlap"`
	ActiveAnnotatorsIDs []int           `db:"active_annotators_ids"`
	InputData           []InputDataRow  `db:"input_data"`
	OutputData          []OutputDataRow `db:"output_data"`
	CreatedAt           time.Time       `db:"created_at"`
}

type InputDataRow struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type OutputDataRow struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	WithAggregation bool   `json:"with_aggregation"`
}

func mapTaskTableToModel(task *TaskTable) model.Task {
	if task == nil {
		return model.Task{}
	}
	taskModel := model.Task{
		ID:                  task.ID,
		ProjectID:           task.ProjectID,
		TargetOverlap:       task.TargetOverlap,
		CurrentOverlap:      task.CurrentOverlap,
		ActiveAnnotatorsIDs: task.ActiveAnnotatorsIDs,
		InputData: lo.Map(task.InputData, func(data InputDataRow, _ int) model.InputData {
			return mapInputDataRowToModel(data)
		}),
		OutputData: lo.Map(task.OutputData, func(data OutputDataRow, _ int) model.OutputData {
			return mapOutputDataRowToModel(data)
		}),
		CreatedAt: task.CreatedAt,
	}
	return taskModel
}

func mapInputDataRowToModel(inputData InputDataRow) model.InputData {
	return model.InputData{
		Name:  inputData.Name,
		Type:  inputData.Type,
		Value: inputData.Value,
	}
}

func mapOutputDataRowToModel(outputData OutputDataRow) model.OutputData {
	return model.OutputData{
		Name:            outputData.Name,
		Type:            outputData.Type,
		Value:           outputData.Value,
		WithAggregation: outputData.WithAggregation,
	}
}

func mapInputDataToDataRow(inputData model.InputData) InputDataRow {
	return InputDataRow{
		Name:  inputData.Name,
		Type:  inputData.Type,
		Value: inputData.Value,
	}
}

func mapOutputDataToDataRow(outputData model.OutputData) OutputDataRow {
	return OutputDataRow{
		Name:            outputData.Name,
		Type:            outputData.Type,
		Value:           outputData.Value,
		WithAggregation: outputData.WithAggregation,
	}
}
