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
	TaskConfig         TaskConfigRow `db:"task_config"`
	TargetOverlap      int           `db:"target_overlap"`
	TasksPerUser       int           `db:"tasks_per_user"`
	AnnotatorTimeLimit time.Duration `db:"annotator_time_limit"`
	CreatedAt          time.Time     `db:"created_at"`
}

type TaskConfigRow struct {
	InputData  []InputDataRow  `json:"input_data"`
	OutputData []OutputDataRow `json:"output_data"`
}

type InputDataRow struct {
	Type       string `json:"type"`
	Name       string `json:"name"`
	IsRequired bool   `json:"is_required"`
}

type OutputDataRow struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	IsRequired      bool   `json:"is_required"`
	WithAggregation bool   `json:"with_aggregation"`
}

func mapProjectTableToModel(projectTable ProjectTable) model.Project {
	return model.Project{
		ID:          projectTable.ID,
		CreatorID:   projectTable.CreatorID,
		Description: projectTable.Description,
		Name:        projectTable.Name,
		Instruction: projectTable.Instruction,
		TaskConfig: model.TaskConfig{
			InputData: lo.Map(projectTable.TaskConfig.InputData,
				func(data InputDataRow, _ int) model.TaskInputData {
					return model.TaskInputData{
						Type:       model.DataType(data.Type),
						Name:       data.Name,
						IsRequired: data.IsRequired,
					}
				}),
			OutputData: lo.Map(projectTable.TaskConfig.OutputData,
				func(data OutputDataRow, _ int) model.TaskOutputData {
					return model.TaskOutputData{
						Type:            model.DataType(data.Type),
						Name:            data.Name,
						IsRequired:      data.IsRequired,
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
