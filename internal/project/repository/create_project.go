package project_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	model "github.com/Flak34/crowd-api/internal/project/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (r *Repository) CreateProject(ctx context.Context, db entrypoint.Database, project model.Project) (int, error) {
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
	projectTable := mapProjectModelToTable(project)
	var id int
	err := pgxscan.Get(ctx, db, &id, query,
		projectTable.CreatorID,
		projectTable.Description,
		projectTable.Name,
		projectTable.Instruction,
		projectTable.Config,
		projectTable.TargetOverlap,
		projectTable.TasksPerUser,
		projectTable.AnnotatorTimeLimit)
	if err != nil {
		return id, errors.Errorf("create project: %s", err.Error())
	}
	return id, nil
}
