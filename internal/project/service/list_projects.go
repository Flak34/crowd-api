package project_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/lib/common"
	model "github.com/Flak34/crowd-api/internal/project/model"
	projectrepo "github.com/Flak34/crowd-api/internal/project/repository"
	"github.com/pkg/errors"
)

const (
	defaultProjectsPageSize  = 10
	defaultProjectsPageNum   = 1
	defaultProjectsSortField = "project.created_at"
)

func (s *Service) ListProjects(ctx context.Context, dto ListProjectsDTO) (projects []model.Project, pagesCount int, err error) {
	selOpts := []common.SelectOption{
		projectrepo.WithCreatorID(dto.Filter.CreatorID),
		projectrepo.WithProjectStatus(dto.Filter.Status),
	}
	db := s.ep.GetDB()
	projectsCount, err := s.projectRepo.CountProjects(ctx, db, selOpts...)
	if err != nil {
		return nil, 0, errors.Wrap(err, "count projects")
	}
	if dto.Sort == nil {
		dto.Sort = &Sort{
			SortField: defaultProjectsSortField,
		}
	}
	if dto.Page == nil {
		dto.Page = &Page{
			PerPage: defaultProjectsPageSize,
			PageNum: defaultProjectsPageNum,
		}
	}
	selOpts = append(selOpts,
		projectrepo.WithSort(dto.Sort.SortField, dto.Sort.Desc),
		projectrepo.WithPage(dto.Page.PerPage, dto.Page.PageNum),
	)
	projects, err = s.projectRepo.ListProjects(ctx, db, selOpts...)
	if err != nil {
		return nil, 0, errors.Wrap(err, "list projects")
	}
	pagesCount = projectsCount / dto.Page.PerPage
	if projectsCount%dto.Page.PerPage != 0 {
		pagesCount++
	}
	return projects, pagesCount, nil
}
