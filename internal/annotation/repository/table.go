package annotation_repository

import (
	model "github.com/Flak34/crowd-api/internal/annotation/model"
	"github.com/samber/lo"
	"time"
)

type AnnotationTable struct {
	TaskID      int       `db:"task_id"`
	AnnotatorID int       `db:"annotator_id"`
	CreatedAt   time.Time `db:"created_at"`
	Data        DataRow   `db:"data"`
}

type DataRow struct {
	OutputData []OutputDataRow `json:"output_data"`
}

type OutputDataRow struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

func mapAnnotationModelToTable(annotation model.Annotation) AnnotationTable {
	return AnnotationTable{
		TaskID:      annotation.TaskID,
		AnnotatorID: annotation.AnnotatorID,
		CreatedAt:   annotation.CreatedAt,
		Data: DataRow{
			OutputData: lo.Map(annotation.OutputData, func(data model.OutputData, _ int) OutputDataRow {
				return mapOutputDataToDataRow(data)
			}),
		},
	}
}

func mapOutputDataToDataRow(outputData model.OutputData) OutputDataRow {
	return OutputDataRow{
		Name:  outputData.Name,
		Value: outputData.Value,
		Type:  outputData.Type,
	}
}
