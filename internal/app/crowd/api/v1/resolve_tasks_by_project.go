package crowdapiv1

import (
	"context"
	ucerrors "github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	"github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ResolveTasksByProject(
	ctx context.Context,
	req *crowd_api_v1.ResolveTasksByProjectRequest,
) (*crowd_api_v1.ResolveTasksByProjectResponse, error) {
	if req.ProjectId == 0 || req.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "Project and user must be specified")
	}
	res, err := i.taskService.ResolveUserTasksByProject(ctx, int(req.ProjectId), int(req.UserId))
	_ = res
	if err != nil {
		return nil, ucerrors.MapToGRPC(err)
	}
	return &crowd_api_v1.ResolveTasksByProjectResponse{
		//Tasks: lo.Map(res, func(task model.Task, _ int) *crowd_api_v1.UserTask {
		//	//return mapTaskModelToProto(task)
		//}),
	}, nil
}
