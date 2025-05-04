package project_repository

import (
	model "github.com/Flak34/crowd-api/internal/project/model"
	"github.com/samber/lo"
	"time"
)

type ProjectTable struct {
	ID                 int           `db:"id"`
	CreatorID          int           `db:"creator_id"`
	Description        string        `db:"description"`
	Name               string        `db:"name"`
	Instruction        string        `db:"instruction"`
	Status             string        `db:"status"`
	Config             ConfigRow     `db:"task_config"`
	TargetOverlap      int           `db:"target_overlap"`
	TasksPerUser       int           `db:"tasks_per_user"`
	AnnotatorTimeLimit time.Duration `db:"annotator_time_limit"`
	CreatedAt          time.Time     `db:"created_at"`
}

type ConfigRow struct {
	InputData        []InputDataRow  `json:"input_data"`
	OutputData       []OutputDataRow `json:"output_data"`
	SerializedLayout string          `json:"serialized_layout"`
}

type InputDataRow struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type OutputDataRow struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	WithAggregation bool   `json:"with_aggregation"`
}

func mapProjectTableToModel(projectTable ProjectTable) model.Project {
	return model.Project{
		ID:          projectTable.ID,
		CreatorID:   projectTable.CreatorID,
		Description: projectTable.Description,
		Name:        projectTable.Name,
		Status:      projectTable.Status,
		Instruction: projectTable.Instruction,
		Config: model.Config{
			InputData: lo.Map(projectTable.Config.InputData,
				func(data InputDataRow, _ int) model.InputData {
					return model.InputData{
						Type: model.DataType(data.Type),
						Name: data.Name,
					}
				}),
			OutputData: lo.Map(projectTable.Config.OutputData,
				func(data OutputDataRow, _ int) model.OutputData {
					return model.OutputData{
						Type:            model.DataType(data.Type),
						Name:            data.Name,
						WithAggregation: data.WithAggregation,
					}
				}),
		},
		TargetOverlap:      projectTable.TargetOverlap,
		TasksPerUser:       projectTable.TasksPerUser,
		CreatedAt:          projectTable.CreatedAt,
		AnnotatorTimeLimit: projectTable.AnnotatorTimeLimit,
	}
}

func mapProjectConfigToConfigRow(config model.Config) ConfigRow {
	return ConfigRow{
		InputData: lo.Map(config.InputData,
			func(data model.InputData, _ int) InputDataRow {
				return InputDataRow{
					Type: string(data.Type),
					Name: data.Name,
				}
			}),
		OutputData: lo.Map(config.OutputData,
			func(data model.OutputData, _ int) OutputDataRow {
				return OutputDataRow{
					Type:            string(data.Type),
					Name:            data.Name,
					WithAggregation: data.WithAggregation,
				}
			}),
		SerializedLayout: config.Layout,
	}
}

type ProjectAnnotatorTable struct {
	ProjectID   int       `db:"project_id"`
	AnnotatorID int       `db:"annotator_id"`
	CreatedAt   time.Time `db:"created_at"`
	TaskIDs     []int     `db:"task_ids"`
}

func mapProjectAnnotatorToModel(projectAnnotator ProjectAnnotatorTable) model.ProjectAnnotator {
	return model.ProjectAnnotator{
		ProjectID:   projectAnnotator.ProjectID,
		AnnotatorID: projectAnnotator.AnnotatorID,
		CreatedAt:   projectAnnotator.CreatedAt,
		TaskIDs:     projectAnnotator.TaskIDs,
	}
}
