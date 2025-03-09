package crowdapiv1

import (
	desc "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	taskmodel "github.com/Flak34/crowd-api/internal/task/model"
)

func mapTaskModelToProto(task taskmodel.Task) *desc.UserTask {
	return &desc.UserTask{
		Id:        int32(task.ID),
		InputData: task.InputData,
	}
}
