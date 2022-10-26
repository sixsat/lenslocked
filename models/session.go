package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session.
	// It will be left empty when look up a session,
	// as we only store the hash of a session token in the database.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// TODO: Implement SessionService.Create
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
