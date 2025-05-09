package crowdapiv1

import (
	"encoding/json"
	desc "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	projectmodel "github.com/Flak34/crowd-api/internal/project/model"
	task_model "github.com/Flak34/crowd-api/internal/task/model"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapProjectModelToProto(project projectmodel.Project) *desc.Project {
	res := &desc.Project{
		Id:                 int32(project.ID),
		CreatorId:          int32(project.CreatorID),
		Name:               project.Name,
		Description:        project.Description,
		Instruction:        project.Instruction,
		TargetOverlap:      int32(project.TargetOverlap),
		TasksPerUser:       int32(project.TasksPerUser),
		CreatedAt:          timestamppb.New(project.CreatedAt),
		AnnotatorTimeLimit: durationpb.New(project.AnnotatorTimeLimit),
		Status:             project.Status,
	}
	configDTO := modelToProjectConfigDTO(project.Config)
	configStr, _ := json.Marshal(configDTO)
	res.TaskConfig = string(configStr)
	return res
}

func mapTaskModelToProto(task task_model.Task) *desc.AnnotatorTask {
	res := &desc.AnnotatorTask{
		Id:             int32(task.ID),
		ProjectId:      int32(task.ProjectID),
		TargetOverlap:  int32(task.TargetOverlap),
		CurrentOverlap: int32(task.CurrentOverlap),
		CreatedAt:      timestamppb.New(task.CreatedAt),
	}
	inputData := make(map[string]TaskInputDataDTO, len(task.InputData))
	for _, data := range task.InputData {
		inputData[data.Name] = TaskInputDataDTO{
			Name:  data.Name,
			Type:  data.Type,
			Value: data.Value,
		}
	}
	outputData := make(map[string]TaskOutputDataDTO, len(task.OutputData))
	for _, data := range task.OutputData {
		outputData[data.Name] = TaskOutputDataDTO{
			Name:            data.Name,
			Type:            data.Type,
			Value:           data.Value,
			WithAggregation: data.WithAggregation,
		}
	}
	inputDataStr, _ := json.Marshal(inputData)
	outputDataStr, _ := json.Marshal(outputData)
	res.InputData = string(inputDataStr)
	res.OutputData = string(outputDataStr)
	return res
}
