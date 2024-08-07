// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: repo.sql

package db

import (
	"context"
)

const createRepo = `-- name: CreateRepo :one
INSERT INTO "repo" (service, owner, name, repo_service_id, webhook_id, service_user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id
`

type CreateRepoParams struct {
	Service       Service
	Owner         string
	Name          string
	RepoServiceID int64
	WebhookID     int64
	ServiceUserID int64
}

func (q *Queries) CreateRepo(ctx context.Context, arg CreateRepoParams) (int64, error) {
	row := q.db.QueryRow(ctx, createRepo,
		arg.Service,
		arg.Owner,
		arg.Name,
		arg.RepoServiceID,
		arg.WebhookID,
		arg.ServiceUserID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteRepo = `-- name: DeleteRepo :exec
DELETE FROM "repo"
WHERE id = $1
`

func (q *Queries) DeleteRepo(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteRepo, id)
	return err
}

const getRepoIDByServiceRepoID = `-- name: GetRepoIDByServiceRepoID :one
SELECT id
FROM "repo"
WHERE service = $1 AND repo_service_id = $2
`

type GetRepoIDByServiceRepoIDParams struct {
	Service       Service
	RepoServiceID int64
}

func (q *Queries) GetRepoIDByServiceRepoID(ctx context.Context, arg GetRepoIDByServiceRepoIDParams) (int64, error) {
	row := q.db.QueryRow(ctx, getRepoIDByServiceRepoID, arg.Service, arg.RepoServiceID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getUserRepos = `-- name: GetUserRepos :many
SELECT r.id, r.service, r.owner, r.name, r.repo_service_id, r.webhook_id, r.service_user_id
FROM "repo" r JOIN "service_user" su ON r.service_user_id = su.id
WHERE su.user_id = $1
`

func (q *Queries) GetUserRepos(ctx context.Context, userID int64) ([]Repo, error) {
	rows, err := q.db.Query(ctx, getUserRepos, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repo
	for rows.Next() {
		var i Repo
		if err := rows.Scan(
			&i.ID,
			&i.Service,
			&i.Owner,
			&i.Name,
			&i.RepoServiceID,
			&i.WebhookID,
			&i.ServiceUserID,
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

const userOwnRepo = `-- name: UserOwnRepo :one
SELECT EXISTS(
    SELECT r.id
    FROM "repo" r JOIN "service_user" su ON r.service_user_id = su.id
    WHERE r.id = $1 AND su.user_id = $2
)
`

type UserOwnRepoParams struct {
	ID     int64
	UserID int64
}

func (q *Queries) UserOwnRepo(ctx context.Context, arg UserOwnRepoParams) (bool, error) {
	row := q.db.QueryRow(ctx, userOwnRepo, arg.ID, arg.UserID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
