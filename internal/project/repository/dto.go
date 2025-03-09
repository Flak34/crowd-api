package project_repository

import "time"

type CreateProjectAnnotatorDTO struct {
	ProjectID   int
	AnnotatorID int
	TaskIDs     []int
}

type CreateProjectDTO struct {
	CreatorID          int
	Description        string
	Name               string
	Instruction        string
	TaskConfig         string
	TargetOverlap      int
	TasksPerUser       int
	AnnotatorTimeLimit time.Duration
}
