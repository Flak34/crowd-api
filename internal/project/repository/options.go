package project_repository

import (
	"github.com/Flak34/crowd-api/internal/lib/common"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5/pgtype"
)

func WithProjectIDs(projectIDs ...int) common.SelectOption {
	return func(sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder {
		if len(projectIDs) == 0 {
			return sb
		}
		return sb.Where(sb.Any("id", "=", pgtype.FlatArray[int](projectIDs)))
	}
}

func WithProjectStatus(status string) common.SelectOption {
	return func(sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder {
		if status == "" {
			return sb
		}
		return sb.Where(sb.Equal("project_status.name", status))
	}
}

func WithCreatorID(creatorID int) common.SelectOption {
	return func(sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder {
		if creatorID == 0 {
			return sb
		}
		return sb.Where(sb.Equal("project.creator_id", creatorID))
	}
}

func WithSort(sortField string, desc bool) common.SelectOption {
	return func(sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder {
		sb.OrderBy(sortField)
		if desc {
			sb.Desc()
		}
		return sb
	}
}

func WithPage(perPage int, pageNum int) common.SelectOption {
	return func(sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder {
		return sb.Offset(perPage * (pageNum - 1)).Limit(perPage)
	}
}
