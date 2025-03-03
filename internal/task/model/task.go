package task_model

import "time"

// Task full model of task
type Task struct {
	ID                  int
	ProjectID           int
	TargetOverlap       int
	CurrentOverlap      int
	ActiveAnnotatorsIDs []int
	InputData           string
	OutputData          string
	MaxAnnotationTime   time.Duration
	CreatedAt           time.Time
}
