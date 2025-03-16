package task_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	projectmodel "github.com/Flak34/crowd-api/internal/project/model"
	projectrepo "github.com/Flak34/crowd-api/internal/project/repository"
	model "github.com/Flak34/crowd-api/internal/task/model"
	taskrepo "github.com/Flak34/crowd-api/internal/task/repository"
	"github.com/jackc/pgx/v5"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
)

type TaskRepo interface {
	ReserveTasks(ctx context.Context, db entrypoint.Database, dto taskrepo.ReserveTasksDTO) ([]model.Task, error)
	ReleaseTasks(ctx context.Context, db entrypoint.Database, userID int, taskIDs ...int) ([]int, error)
	ListUserProjectTasks(ctx context.Context, db entrypoint.Database, projectID int, userID int) ([]model.Task, error)
}

type ProjectRepo interface {
	GetProject(ctx context.Context, db entrypoint.Database, projectID int) (projectmodel.Project, error)
	CreateProjectAnnotator(ctx context.Context, db entrypoint.Database, dto projectrepo.CreateProjectAnnotatorDTO) error
	GetProjectAnnotator(ctx context.Context, db entrypoint.Database, projectID int, userID int) (projectmodel.ProjectAnnotator, error)
}

type PgqClient[TTx any] interface {
	InsertTx(ctx context.Context, tx TTx, args river.JobArgs, opts *river.InsertOpts) (*rivertype.JobInsertResult, error)
	Insert(ctx context.Context, args river.JobArgs, opts *river.InsertOpts) (*rivertype.JobInsertResult, error)
}

type Service struct {
	taskRepo    TaskRepo
	projectRepo ProjectRepo
	pgqClient   PgqClient[pgx.Tx]
	ep          entrypoint.Entrypoint
}

func New(ep entrypoint.Entrypoint, taskRepo TaskRepo, projectRepo ProjectRepo, pgqClient PgqClient[pgx.Tx]) *Service {
	return &Service{
		taskRepo:    taskRepo,
		projectRepo: projectRepo,
		pgqClient:   pgqClient,
		ep:          ep,
	}
}
