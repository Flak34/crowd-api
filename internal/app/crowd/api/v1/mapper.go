package crowdapiv1

import (
	desc "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	projectmodel "github.com/Flak34/crowd-api/internal/project/model"
	taskmodel "github.com/Flak34/crowd-api/internal/task/model"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapTaskModelToProto(task taskmodel.Task) *desc.UserTask {
	return &desc.UserTask{
		Id:        int32(task.ID),
		InputData: task.InputData,
	}
}

func mapProjectModelToProto(project projectmodel.Project) *desc.Project {
	return &desc.Project{
		Id:                 int32(project.ID),
		CreatorId:          int32(project.CreatorID),
		Name:               project.Name,
		Description:        project.Description,
		Instruction:        project.Instruction,
		TargetOverlap:      int32(project.TargetOverlap),
		TasksPerUser:       int32(project.TasksPerUser),
		CreatedAt:          timestamppb.New(project.CreatedAt),
		AnnotatorTimeLimit: durationpb.New(project.AnnotatorTimeLimit),
	}
}
