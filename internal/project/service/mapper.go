package project_service

import (
	"encoding/json"
	model "github.com/Flak34/crowd-api/internal/project/model"
	projectrepo "github.com/Flak34/crowd-api/internal/project/repository"
)

func mapProjectModelToDTO(project model.Project) projectrepo.CreateProjectDTO {
	taskConfigStr, _ := json.Marshal(project.TaskConfig)
	return projectrepo.CreateProjectDTO{
		CreatorID:          project.CreatorID,
		Description:        project.Description,
		Name:               project.Name,
		Instruction:        project.Instruction,
		TaskConfig:         string(taskConfigStr),
		TargetOverlap:      project.TargetOverlap,
		TasksPerUser:       project.TasksPerUser,
		AnnotatorTimeLimit: project.AnnotatorTimeLimit,
	}
}
