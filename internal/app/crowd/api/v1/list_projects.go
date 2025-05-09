package crowdapiv1

import (
	"context"
	crowd_api_v1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	model "github.com/Flak34/crowd-api/internal/project/model"
	project_service "github.com/Flak34/crowd-api/internal/project/service"
	"github.com/samber/lo"
)

func (i *Implementation) ListProjects(ctx context.Context, req *crowd_api_v1.ListProjectsRequest) (*crowd_api_v1.ListProjectsResponse, error) {
	dto := project_service.ListProjectsDTO{
		Filter: project_service.ProjectsFilter{
			CreatorID: int(req.GetFilter().GetCreatorId()),
			Status:    req.GetFilter().GetStatus(),
		},
	}
	if req.GetPage() != nil {
		dto.Page = &project_service.Page{
			PerPage: int(req.GetPage().GetPerPage()),
			PageNum: int(req.GetPage().GetPageNum()),
		}
	}
	if req.GetSort() != nil {
		dto.Sort = &project_service.Sort{
			SortField: req.GetSort().GetSortField(),
			Desc:      req.GetSort().GetDesc(),
		}
	}

	projects, pagesCount, err := i.projectService.ListProjects(ctx, dto)
	if err != nil {
		return nil, err
	}

	return &crowd_api_v1.ListProjectsResponse{
		Projects: lo.Map(projects, func(project model.Project, _ int) *crowd_api_v1.Project {
			return mapProjectModelToProto(project)
		}),
		Page: &crowd_api_v1.PageResponse{PagesCount: int32(pagesCount)},
	}, nil
}
