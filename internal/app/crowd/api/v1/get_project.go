package crowdapiv1

import (
	"context"
	uscerrors "github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	desc "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
)

func (i *Implementation) GetProject(ctx context.Context, req *desc.GetProjectRequest) (*desc.GetProjectResponse, error) {
	project, err := i.projectService.GetProject(ctx, int(req.GetId()))
	if err != nil {
		return nil, uscerrors.MapToGRPC(err)
	}
	return &desc.GetProjectResponse{
		Project: mapProjectModelToProto(project),
	}, nil
}
