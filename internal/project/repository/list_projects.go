package project_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/lib/common"
	model "github.com/Flak34/crowd-api/internal/project/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/huandu/go-sqlbuilder"
	"github.com/samber/lo"
)

func (r *Repository) ListProjects(ctx context.Context, db entrypoint.Database, opts ...common.SelectOption) ([]model.Project, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.
		From("project").
		Select(
			"project.id",
			"project.creator_id",
			"project.description",
			"project.name",
			"project.instruction",
			"project.task_config",
			"project.target_overlap",
			"project.tasks_per_user",
			"project.created_at",
			"project.annotator_time_limit",
			"project_status.name AS status_name").
		JoinWithOption(sqlbuilder.LeftJoin, "project_status", "project.status_id = project_status.id")
	for _, opt := range opts {
		sb = opt(sb)
	}
	query, args := sb.Build()
	var projects []*ProjectTable
	err := pgxscan.Select(ctx, db, &projects, query, args...)
	if err != nil {
		return []model.Project{}, err
	}
	return lo.FilterMap(projects, func(p *ProjectTable, _ int) (model.Project, bool) {
		if p == nil {
			return model.Project{}, false
		}
		return mapProjectTableToModel(*p), true
	}), nil
}
