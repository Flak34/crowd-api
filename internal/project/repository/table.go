package project_repository

import (
	model "github.com/Flak34/crowd-api/internal/project/model"
	"time"
)

type ProjectTable struct {
	ID            int       `db:"id"`
	CreatorID     int       `db:"creator_id"`
	Description   string    `db:"description"`
	TaskConfig    string    `db:"task_config"`
	TargetOverlap int       `db:"target_overlap"`
	TasksPerUser  int       `db:"tasks_per_user"`
	CreatedAt     time.Time `db:"created_at"`
}

func mapProjectTableToModel(projectTable ProjectTable) model.Project {
	return model.Project{
		ID:            projectTable.ID,
		CreatorID:     projectTable.CreatorID,
		Description:   projectTable.Description,
		TaskConfig:    projectTable.TaskConfig,
		TargetOverlap: projectTable.TargetOverlap,
		TasksPerUser:  projectTable.TasksPerUser,
		CreatedAt:     projectTable.CreatedAt,
	}
}

type ProjectAnnotatorTable struct {
	ProjectID   int   `db:"project_id"`
	AnnotatorID int   `db:"annotator_id"`
	TaskIDs     []int `db:"task_ids"`
}

func mapProjectAnnotatorToModel(projectAnnotator ProjectAnnotatorTable) model.ProjectAnnotator {
	return model.ProjectAnnotator{
		ProjectID:   projectAnnotator.ProjectID,
		AnnotatorID: projectAnnotator.AnnotatorID,
		TaskIDs:     projectAnnotator.TaskIDs,
	}
}
