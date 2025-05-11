package crowdapiv1

import (
	"context"
	"encoding/json"
	model "github.com/Flak34/crowd-api/internal/annotation/model"
	desc "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	"github.com/samber/lo"
)

func (i *Implementation) CreateAnnotations(ctx context.Context, req *desc.CreateAnnotationsRequest) (*desc.CreateAnnotationsResponse, error) {
	annotations := make([]model.Annotation, 0, len(req.GetAnnotations()))
	for _, annotation := range req.GetAnnotations() {
		var data []AnnotationOutputDataDTO
		json.Unmarshal([]byte(annotation.GetOutputData()), &data)
		annotations = append(annotations, model.Annotation{
			TaskID: int(annotation.GetTaskId()),
			OutputData: lo.Map(data, func(outputData AnnotationOutputDataDTO, _ int) model.OutputData {
				return model.OutputData{
					Type:  outputData.Type,
					Name:  outputData.Name,
					Value: outputData.Value,
				}
			}),
		})
	}
	err := i.annotationsService.CreateAnnotations(ctx, annotations...)
	if err != nil {
		return nil, err
	}
	return &desc.CreateAnnotationsResponse{}, nil
}
