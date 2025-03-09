package annotation_model

import "time"

type Annotation struct {
	TaskID      int
	AnnotatorId int
	CreatedAt   time.Time
	OutputData  AnnotationOutputData
}

type AnnotationOutputData struct {
}

type DataType string

type InputData struct {
	Type       DataType `json:"type"`
	Name       string   `json:"name"`
	IsRequired bool     `json:"is_required"`
}

type OutputData struct {
	Type            DataType `json:"type"`
	Name            string   `json:"name"`
	IsRequired      bool     `json:"is_required"`
	WithAggregation bool     `json:"with_aggregation"`
}
