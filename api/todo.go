package api

import (
	db "MasonPhan2110/Todo_Go_GPT/db/sqlc"
	"MasonPhan2110/Todo_Go_GPT/utils"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type createTaskRequest struct {
	UserID      int64          `json:"userid" binding:"required,min=1"`
	Name        string         `json:"name" binding:"required"`
	Description sql.NullString `json:"description" binding:"required"`
	Deadline    time.Time      `json:"deadline" binding:"required"`
}

type taskResponse struct {
	ID          int64          `json:"id"`
	UserID      int64          `json:"user_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Status      bool           `json:"status"`
	Deadline    time.Time      `json:"deadline"`
	UpdateAt    time.Time      `json:"update_at"`
	CreatedAt   time.Time      `json:"created_at"`
}

func newTaskResponse(task db.Todo) taskResponse {
	return taskResponse{
		ID:          task.ID,
		UserID:      task.UserID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		UpdateAt:    task.UpdateAt,
		Deadline:    task.Deadline,
		CreatedAt:   task.CreatedAt,
	}
}

func (server *Server) CreateTask(ctx *gin.Context) {
	var req createTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	arg := db.CreateTaskParams{
		UserID:      req.UserID,
		Name:        req.Name,
		Description: req.Description,
		Deadline:    req.Deadline,
	}

	task, err := db.DBStore.CreateTask(ctx, arg)

	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	rsp := newTaskResponse(task)
	ctx.JSON(http.StatusOK, rsp)
}

type getTaskRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetTask(ctx *gin.Context) {
	var req getTaskRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	task, err := db.DBStore.GetTask(ctx, req.ID)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	rsp := newTaskResponse(task)
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) UpdateTask(ctx *gin.Context) {}

func (server *Server) DeleteTask(ctx *gin.Context) {}
