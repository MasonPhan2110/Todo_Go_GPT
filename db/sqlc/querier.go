// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserForUpdate(ctx context.Context, username string) (User, error)
	UpdateUserHashedPassword(ctx context.Context, arg UpdateUserHashedPasswordParams) (User, error)
}

var _ Querier = (*Queries)(nil)