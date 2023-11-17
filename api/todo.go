package api

import (
	db "MasonPhan2110/Todo_Go_GPT/db/sqlc"
	"MasonPhan2110/Todo_Go_GPT/middleware"
	"MasonPhan2110/Todo_Go_GPT/pkg/token"
	"MasonPhan2110/Todo_Go_GPT/utils"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type createTaskRequest struct {
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

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	arg := db.CreateTaskParams{
		UserID:      authPayload.UserId,
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

func (server *Server) GetTask(ctx *gin.Context) {}

type listTasksRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListTasks(ctx *gin.Context) {
	var req listTasksRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	arg := db.ListTasksParams{
		UserID: authPayload.UserId,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
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

type updateTaskRequest struct {
	ID          int64     `json:"id" binding:"required,min=1"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Deadline    time.Time `json:"deadline"`
}

func (server *Server) UpdateTask(ctx *gin.Context) {
	var req updateTaskRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	if validUser, err := server.ValidUser(ctx, authPayload.UserId, req.ID); !validUser && err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	args := db.UpdateTaskParams{
		Name: sql.NullString{
			String: req.Name,
			Valid:  req.Name != "",
		},
		Description: sql.NullString{
			String: req.Description,
			Valid:  req.Description != "",
		},
		Deadline: sql.NullTime{
			Time:  req.Deadline,
			Valid: !req.Deadline.IsZero(),
		},
		ID: req.ID,
	}

	if req.Status != "" {
		value, err := strconv.ParseBool(req.Status)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}
		args.Status = sql.NullBool{
			Bool:  value,
			Valid: true,
		}
	}

	task, err := db.DBStore.UpdateTask(ctx, args)
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

type deleteTaskRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) DeleteTask(ctx *gin.Context) {
	var req deleteTaskRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	if validUser, err := server.ValidUser(ctx, authPayload.UserId, req.ID); !validUser && err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	err := db.DBStore.DeleteTask(ctx, req.ID)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (server *Server) ValidUser(ctx *gin.Context, userID int64, ID int64) (bool, error) {
	task, err := db.DBStore.GetTask(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return false, err
	}

	if task.UserID != userID {
		err := fmt.Errorf("Unauthorize")
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return false, err
	}

	return true, nil

}
