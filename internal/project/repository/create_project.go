package project_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (r *Repository) CreateProject(ctx context.Context, db entrypoint.Database, dto CreateProjectDTO) (int, error) {
	var query = `
		INSERT 
		INTO project(
		             creator_id, 
		             description, 
		             name, 
		             instruction, 
		             task_config, 
		             target_overlap, 
		             tasks_per_user, 
		             annotator_time_limit)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`
	var id int
	err := pgxscan.Get(ctx, db, &id, query,
		dto.CreatorID,
		dto.Description,
		dto.Name,
		dto.Instruction,
		dto.TaskConfig,
		dto.TargetOverlap,
		dto.TasksPerUser,
		dto.AnnotatorTimeLimit)
	if err != nil {
		return id, errors.Errorf("create project: %s", err.Error())
	}
	return id, nil
}
