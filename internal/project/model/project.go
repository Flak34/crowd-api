package project_model

import (
	"github.com/pkg/errors"
	"time"
)

type Project struct {
	ID                 int
	CreatorID          int
	Description        string
	Name               string
	Instruction        string
	Config             Config
	TargetOverlap      int
	TasksPerUser       int
	AnnotatorTimeLimit time.Duration
	CreatedAt          time.Time
}

type Config struct {
	InputData  []InputData
	OutputData []OutputData
}

type DataType string

type InputData struct {
	Type       DataType
	Name       string
	IsRequired bool
}

type OutputData struct {
	Type            DataType
	Name            string
	IsRequired      bool
	WithAggregation bool
}

func (c Config) Validate() error {
	for _, data := range c.InputData {
		if !data.Type.IsValid() {
			return errors.Errorf("invalid input data type: %s", data.Type)
		}
	}
	for _, data := range c.OutputData {
		if !data.Type.IsValid() {
			return errors.Errorf("invalid output data type: %s", data.Type)
		}
	}
	return nil
}

func (t DataType) IsValid() bool {
	switch t {
	case "string", "int", "[]int", "[]string", "link", "bool":
		return true
	default:
		return false
	}
}
