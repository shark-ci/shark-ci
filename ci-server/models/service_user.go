package models

import (
	"time"

	"golang.org/x/oauth2"
)

// TODO: All nullable fields should be pointers.
// TODO: Tokens shouldn't be serialized to JSON.
type ServiceUser struct {
	ID           int64     `json:"id"`
	Service      string    `json:"service"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	TokenExpire  time.Time `json:"token_expire"`
	UserID       int64     `json:"user_id"`
}

func (su ServiceUser) Token() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  su.AccessToken,
		RefreshToken: su.RefreshToken,
		TokenType:    su.TokenType,
		Expiry:       su.TokenExpire,
	}
}
