package api

import (
	db "MasonPhan2110/Todo_Go_GPT/db/sqlc"
	"MasonPhan2110/Todo_Go_GPT/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type createTaskRequest struct {
	UserID      int64     `json:"user_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
}

type createTaskResponse struct {
	UserID      int64     `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
}

func newTaskResponse(task db.Todo) createTaskResponse {
	return createTaskResponse{
		UserID:      task.UserID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		Deadline:    task.Deadline,
		CreatedAt:   task.CreatedAt,
	}
}

// CreateTask godoc
//
//	@Summary		Create new Task to Todo
//	@Tags			todo
//	@Accept			json
//	@Produce		json
//	@Param			request body createTaskRequest true "query params"
//	@Failure		400	{object}	utils.HTTPError
//	@Failure		404	{object}	utils.HTTPError
//	@Failure		500	{object}	utils.HTTPError
//	@Success		200	{object}	createTaskResponse
//	@Router			/api/v1/todo/create [post]
func (server *Server) CreateTask(ctx *gin.Context) {
	var req createTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	arg := db.CreateTaskParams{
		UserID:      req.UserID,
		Name:        req.Name,
		Description: req.Description,
		Status:      false,
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
	UserID int64 `json:"user_id" binding:"required"`
	Limit  int32 `json:"limit" binding:"required"`
	Offset int32 `json:"offset" binding:"required"`
}

type getTaskResponse struct {
	UserID      int64     `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
}

func (server *Server) GetTasks(ctx *gin.Context) {
	var req getTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	arg := db.ListTasksParams{
		UserID: req.UserID,
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	listTasks, err := db.DBStore.ListTasks(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, listTasks)
}

func (server *Server) UpdateTask(ctx *gin.Context) {}

func (server *Server) DeleteTask(ctx *gin.Context) {}
