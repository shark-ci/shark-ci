// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Service string

const (
	ServiceGitHub Service = "GitHub"
	ServiceGitLab Service = "GitLab"
)

func (e *Service) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Service(s)
	case string:
		*e = Service(s)
	default:
		return fmt.Errorf("unsupported scan type for Service: %T", src)
	}
	return nil
}

type NullService struct {
	Service Service
	Valid   bool // Valid is true if Service is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullService) Scan(value interface{}) error {
	if value == nil {
		ns.Service, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Service.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullService) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Service), nil
}

type Oauth2State struct {
	State  pgtype.UUID
	Expire pgtype.Timestamp
}

type Pipeline struct {
	ID         int64
	Url        pgtype.Text
	Status     string
	CloneUrl   string
	CommitSha  string
	StartedAt  pgtype.Timestamp
	FinishedAt pgtype.Timestamp
	RepoID     pgtype.Int8
}

type Repo struct {
	ID            int64
	Service       Service
	Owner         string
	Name          string
	RepoServiceID int64
	WebhookID     pgtype.Int8
	ServiceUserID pgtype.Int8
}

type ServiceUser struct {
	ID           int64
	Service      Service
	Username     string
	Email        string
	AccessToken  string
	RefreshToken pgtype.Text
	TokenType    string
	TokenExpire  pgtype.Timestamp
	UserID       pgtype.Int8
}

type User struct {
	ID       int64
	Username string
	Email    string
}
