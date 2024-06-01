// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: service_user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createServiceUser = `-- name: CreateServiceUser :one
INSERT INTO "service_user" (service, username, email, access_token, refresh_token, token_type, token_expire, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id
`

type CreateServiceUserParams struct {
	Service      Service
	Username     string
	Email        string
	AccessToken  string
	RefreshToken pgtype.Text
	TokenType    string
	TokenExpire  pgtype.Timestamp
	UserID       int64
}

func (q *Queries) CreateServiceUser(ctx context.Context, arg CreateServiceUserParams) (int64, error) {
	row := q.db.QueryRow(ctx, createServiceUser,
		arg.Service,
		arg.Username,
		arg.Email,
		arg.AccessToken,
		arg.RefreshToken,
		arg.TokenType,
		arg.TokenExpire,
		arg.UserID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getServiceUserByUserID = `-- name: GetServiceUserByUserID :one
SELECT id, username, username, email, access_token, refresh_token, token_type, token_expire
FROM "service_user"
WHERE user_id = $1 AND service = $2
`

type GetServiceUserByUserIDParams struct {
	UserID  int64
	Service Service
}

type GetServiceUserByUserIDRow struct {
	ID           int64
	Username     string
	Username_2   string
	Email        string
	AccessToken  string
	RefreshToken pgtype.Text
	TokenType    string
	TokenExpire  pgtype.Timestamp
}

func (q *Queries) GetServiceUserByUserID(ctx context.Context, arg GetServiceUserByUserIDParams) (GetServiceUserByUserIDRow, error) {
	row := q.db.QueryRow(ctx, getServiceUserByUserID, arg.UserID, arg.Service)
	var i GetServiceUserByUserIDRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Username_2,
		&i.Email,
		&i.AccessToken,
		&i.RefreshToken,
		&i.TokenType,
		&i.TokenExpire,
	)
	return i, err
}

const getUserIDByServiceUser = `-- name: GetUserIDByServiceUser :one
SELECT user_id
FROM "service_user"
WHERE service = $1 AND username = $2
`

type GetUserIDByServiceUserParams struct {
	Service  Service
	Username string
}

func (q *Queries) GetUserIDByServiceUser(ctx context.Context, arg GetUserIDByServiceUserParams) (int64, error) {
	row := q.db.QueryRow(ctx, getUserIDByServiceUser, arg.Service, arg.Username)
	var user_id int64
	err := row.Scan(&user_id)
	return user_id, err
}
