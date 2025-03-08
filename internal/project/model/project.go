package project_model

import (
	"time"
)

type Project struct {
	ID                 int
	CreatorID          int
	Description        string
	TaskConfig         string
	TargetOverlap      int
	TasksPerUser       int
	AnnotatorTimeLimit time.Duration
	CreatedAt          time.Time
}

type ProjectAnnotator struct {
	ProjectID   int
	AnnotatorID int
	TaskIDs     []int
}
