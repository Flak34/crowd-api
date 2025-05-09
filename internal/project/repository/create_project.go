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
		             annotator_time_limit,
		             status_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, (SELECT id FROM project_status WHERE name = $9))
		RETURNING id`
	configRow := mapProjectConfigToConfigRow(project.Config)
	var id int
	err := pgxscan.Get(ctx, db, &id, query,
		project.CreatorID,
		project.Description,
		project.Name,
		project.Instruction,
		configRow,
		project.TargetOverlap,
		project.TasksPerUser,
		project.AnnotatorTimeLimit,
		project.Status)
	if err != nil {
		return id, errors.Errorf("create project: %s", err.Error())
	}
	return id, nil
}
