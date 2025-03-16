package annotation_model

import "time"

type Annotation struct {
	TaskID      int
	AnnotatorID int
	CreatedAt   time.Time
	OutputData  []AnnotationOutputData
}

type AnnotationOutputData struct {
	Name string
	Data string
}
