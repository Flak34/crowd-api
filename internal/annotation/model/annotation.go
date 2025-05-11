package annotation_model

import "time"

type Annotation struct {
	TaskID      int
	AnnotatorID int
	CreatedAt   time.Time
	OutputData  []OutputData
}

type OutputData struct {
	Name  string
	Value string
	Type  string
}
