package project_model

import (
	annotationmodel "github.com/Flak34/crowd-api/internal/annotation/model"
	"github.com/pkg/errors"
	"time"
)

type Project struct {
	ID                 int
	CreatorID          int
	Description        string
	Name               string
	Status             string
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
	Layout     string
}

type DataType string

type InputData struct {
	Type DataType
	Name string
}

type OutputData struct {
	Type            DataType
	Name            string
	WithAggregation bool
}

func (p Project) ValidateAnnotation(annotation annotationmodel.Annotation) error {
	return nil
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
