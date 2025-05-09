package task_model

import "time"

// Task full model of task
type Task struct {
	ID                  int
	ProjectID           int
	TargetOverlap       int
	CurrentOverlap      int
	ActiveAnnotatorsIDs []int
	InputData           []InputData
	OutputData          []OutputData
	CreatedAt           time.Time
}

type InputData struct {
	Name  string
	Type  string
	Value string
}

type OutputData struct {
	Name            string
	Type            string
	Value           string
	WithAggregation bool
}

func GetDefaultDataValue(dataType string) string {
	switch dataType {
	case "int":
		return "0"
	case "string":
		return ""
	case "bool":
		return "false"
	}
	return ""
}
