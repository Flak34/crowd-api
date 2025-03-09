package crowdapiv1

import (
	"context"
	uscerrors "github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	desc "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateProject(ctx context.Context, req *desc.CreateProjectRequest) (*desc.CreateProjectResponse, error) {
	dto, err := protoToProjectDTO(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Create project req is invalid: %v", err)
	}
	projectModel := projectDTOToModel(dto)
	id, err := i.projectService.CreateProject(ctx, projectModel)
	if err != nil {
		return nil, uscerrors.MapToGRPC(err)
	}
	return &desc.CreateProjectResponse{Id: int32(id)}, nil
}
