package task_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	task_model "github.com/Flak34/crowd-api/internal/task/model"
	"github.com/huandu/go-sqlbuilder"
	"github.com/samber/lo"
)

func (r *Repository) CreateTasks(ctx context.Context, db entrypoint.Database, tasks ...task_model.Task) error {
	insB := sqlbuilder.PostgreSQL.NewInsertBuilder()
	insB.InsertInto("task").
		Cols(
			"project_id",
			"target_overlap",
			"current_overlap",
			"input_data",
			"output_data")
	for _, task := range tasks {
		inputData := lo.Map(task.InputData, func(data task_model.InputData, _ int) InputDataRow {
			return mapInputDataToDataRow(data)
		})
		outputData := lo.Map(task.OutputData, func(data task_model.OutputData, _ int) OutputDataRow {
			return mapOutputDataToDataRow(data)
		})
		insB.Values(task.ProjectID, task.TargetOverlap, task.CurrentOverlap, inputData, outputData)
	}
	query, args := insB.Build()
	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}
