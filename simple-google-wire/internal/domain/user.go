package domain

import (
	"context"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserEntity struct {
	ID       string
	Username string
	Password string
	Email    string
}

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*UserEntity, error)
}

type UserService interface {
	FindByUsername(ctx context.Context, username string) (*User, error)
}

type UserHandler interface {
	FindByUsername() http.HandlerFunc
}
