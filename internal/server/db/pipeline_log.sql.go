// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: pipeline_log.sql

package db

import (
	"context"
)

const createPipelineLog = `-- name: CreatePipelineLog :one
INSERT INTO "pipeline_log" ("order", "cmd", "output", "exit_code", "pipeline_id")
VALUES ($1, $2, $3, $4, $5)
RETURNING "id"
`

type CreatePipelineLogParams struct {
	Order      int32
	Cmd        string
	Output     string
	ExitCode   int32
	PipelineID int64
}

func (q *Queries) CreatePipelineLog(ctx context.Context, arg CreatePipelineLogParams) (int64, error) {
	row := q.db.QueryRow(ctx, createPipelineLog,
		arg.Order,
		arg.Cmd,
		arg.Output,
		arg.ExitCode,
		arg.PipelineID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getPipelineLogs = `-- name: GetPipelineLogs :many
SELECT "order", "cmd", "output", "exit_code"
FROM "pipeline_log"
WHERE "pipeline_id" = $1
ORDER BY "order"
`

type GetPipelineLogsRow struct {
	Order    int32
	Cmd      string
	Output   string
	ExitCode int32
}

func (q *Queries) GetPipelineLogs(ctx context.Context, pipelineID int64) ([]GetPipelineLogsRow, error) {
	rows, err := q.db.Query(ctx, getPipelineLogs, pipelineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPipelineLogsRow
	for rows.Next() {
		var i GetPipelineLogsRow
		if err := rows.Scan(
			&i.Order,
			&i.Cmd,
			&i.Output,
			&i.ExitCode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
