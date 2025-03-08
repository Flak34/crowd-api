-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS task (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL,
    target_overlap INTEGER NOT NULL,
    current_overlap INTEGER NOT  NULL,
    active_annotators_ids INTEGER[],
    input_data JSONB NOT NULL,
    output_data JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    annotated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
CREATE INDEX project_id_idx ON task(project_id);

CREATE TABLE IF NOT EXISTS project (
    id SERIAL PRIMARY KEY,
    creator_id INTEGER NOT NULL,
    description TEXT NOT NULL,
    task_config JSONB NOT NULL,
    target_overlap INTEGER NOT NULL,
    tasks_per_user INTEGER NOT NULL,
    annotator_time_limit INTERVAL NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc')
);

CREATE TABLE IF NOT EXISTS task_annotation (
    task_id INTEGER NOT NULL,
    annotator_id TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    output_data JSONB,
    PRIMARY KEY (task_id, annotator_id)
);

CREATE TABLE IF NOT EXISTS project_annotator (
  project_id INTEGER NOT NULL,
  annotator_id INTEGER NOT NULL,
  task_ids INTEGER[] NOT NULL ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
  PRIMARY KEY (project_id, annotator_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS task, project, task_annotation, project_annotator;
-- +goose StatementEnd
