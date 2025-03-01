-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS task (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL,
    target_overlap INTEGER NOT NULL,
    current_overlap INTEGER NOT  NULL,
    active_annotators_ids TEXT[],
    input_data JSONB NOT NULL,
    output_data JSONB,
    time_to_annotate INTERVAL NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    annotated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
CREATE INDEX project_id_idx ON task(project_id);

CREATE TABLE IF NOT EXISTS project (
    id SERIAL PRIMARY KEY,
    creator_id TEXT,
    description TEXT,
    target_overlap INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc')
);

CREATE TABLE IF NOT EXISTS task_annotation (
    task_id INTEGER,
    annotator_id TEXT NOT NULL,
    started_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    deadline TIMESTAMPTZ,
    finished_at TIMESTAMPTZ,
    output_data JSONB,
    PRIMARY KEY (task_id, annotator_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS task, project, task_annotation;
-- +goose StatementEnd
