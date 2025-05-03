package project_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/lib/common"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/huandu/go-sqlbuilder"
)

func (r *Repository) CountProjects(ctx context.Context, db entrypoint.Database, opts ...common.SelectOption) (int, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.
		From("project").
		Select("COUNT(*)").
		JoinWithOption(sqlbuilder.LeftJoin, "project_status", "project.status_id = project_status.id")
	for _, opt := range opts {
		sb = opt(sb)
	}
	query, args := sb.Build()
	var count int
	err := pgxscan.Get(ctx, db, &count, query, args...)
	if err != nil {
		return count, err
	}
	return count, nil
}
