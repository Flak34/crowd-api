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
	Name string `json:"name"`
	Data string `json:"data"`
}

func mapAnnotationModelToTable(annotation model.Annotation) AnnotationTable {
	return AnnotationTable{
		TaskID:      annotation.TaskID,
		AnnotatorID: annotation.AnnotatorID,
		CreatedAt:   annotation.CreatedAt,
		Data: DataRow{
			OutputData: lo.Map(annotation.OutputData, func(data model.AnnotationOutputData, _ int) OutputDataRow {
				return OutputDataRow{
					Name: data.Name,
					Data: data.Data,
				}
			}),
		},
	}
}

func mapAnnotationTableToModel(annotation AnnotationTable) model.Annotation {
	return model.Annotation{
		TaskID:      annotation.TaskID,
		AnnotatorID: annotation.AnnotatorID,
		CreatedAt:   annotation.CreatedAt,
		OutputData: lo.Map(annotation.Data.OutputData, func(data OutputDataRow, _ int) model.AnnotationOutputData {
			return model.AnnotationOutputData{
				Name: data.Name,
				Data: data.Data,
			}
		}),
	}
}
