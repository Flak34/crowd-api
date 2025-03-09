package project_model

import (
	"time"
)

type Project struct {
	ID                 int
	CreatorID          int
	Description        string
	Name               string
	Instruction        string
	TaskConfig         TaskConfig
	TargetOverlap      int
	TasksPerUser       int
	AnnotatorTimeLimit time.Duration
	CreatedAt          time.Time
}

type ProjectAnnotator struct {
	ProjectID   int
	AnnotatorID int
	TaskIDs     []int
	CreatedAt   time.Time
}

type TaskConfig struct {
	InputData  []TaskInputData  `json:"input_data"`
	OutputData []TaskOutputData `json:"output_data"`
}

type DataType string

type TaskInputData struct {
	Type       DataType `json:"type"`
	Name       string   `json:"name"`
	IsRequired bool     `json:"is_required"`
}

type TaskOutputData struct {
	Type            DataType `json:"type"`
	Name            string   `json:"name"`
	IsRequired      bool     `json:"is_required"`
	WithAggregation bool     `json:"with_aggregation"`
}

func (t DataType) IsValid() bool {
	switch t {
	case "string", "int", "[]int", "[]string", "link", "bool":
		return true
	default:
		return false
	}
}
