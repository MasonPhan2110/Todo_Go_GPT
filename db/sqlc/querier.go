// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTask(ctx context.Context, arg CreateTaskParams) (Todo, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	DeleteTask(ctx context.Context, id int64) error
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTask(ctx context.Context, id int64) (Todo, error)
	GetTaskForUpdate(ctx context.Context, id int64) (Todo, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserForUpdate(ctx context.Context, username string) (User, error)
	ListTasks(ctx context.Context, arg ListTasksParams) ([]Todo, error)
	UpdateTask(ctx context.Context, arg UpdateTaskParams) (Todo, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
