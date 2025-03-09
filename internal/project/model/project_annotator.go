package project_model

import "time"

type ProjectAnnotator struct {
	ProjectID   int
	AnnotatorID int
	TaskIDs     []int
	CreatedAt   time.Time
}
