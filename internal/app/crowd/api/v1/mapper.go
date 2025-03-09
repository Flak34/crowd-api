package crowdapiv1

import (
	"encoding/json"
	desc "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	projectmodel "github.com/Flak34/crowd-api/internal/project/model"
	model "github.com/Flak34/crowd-api/internal/task/model"
	"github.com/pkg/errors"
)

func mapTaskModelToProto(task model.Task) *desc.UserTask {
	return &desc.UserTask{
		Id:        int32(task.ID),
		InputData: task.InputData,
	}
}

func mapProtoProjectRequestToModel(req *desc.CreateProjectRequest) (projectmodel.Project, error) {
	if req == nil {
		return projectmodel.Project{}, nil
	}
	var taskConfig projectmodel.TaskConfig
	err := json.Unmarshal([]byte(req.TaskConfig), &taskConfig)
	if err != nil {
		return projectmodel.Project{}, errors.Errorf("unmarshal task config: %s", err.Error())
	}
	return projectmodel.Project{
		CreatorID:          int(req.CreatorId),
		Description:        req.Description,
		Name:               req.Name,
		Instruction:        req.Instruction,
		TaskConfig:         taskConfig,
		TargetOverlap:      int(req.TargetOverlap),
		TasksPerUser:       int(req.TasksPerUser),
		AnnotatorTimeLimit: req.AnnotatorTimeLimit.AsDuration(),
	}, nil
}
