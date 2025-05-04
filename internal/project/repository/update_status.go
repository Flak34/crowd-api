package project_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
)

func (r *Repository) UpdateStatus(ctx context.Context, db entrypoint.Database, projectID int, newStatus string) error {
	query := `UPDATE project SET status_id = (SELECT id FROM project_status WHERE name = $1) WHERE id = $2`
	rows, err := db.Query(ctx, query, newStatus, projectID)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}
