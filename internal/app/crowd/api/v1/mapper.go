package crowdapiv1

import (
	crowd_api_v1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	model "github.com/Flak34/crowd-api/internal/task/model"
)

func mapTaskToProto(task model.Task) *crowd_api_v1.UserTask {
	return &crowd_api_v1.UserTask{
		Id:        int32(task.ID),
		InputData: task.InputData,
	}
}
