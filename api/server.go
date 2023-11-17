package api

import (
	db "MasonPhan2110/Todo_Go_GPT/db/sqlc"
	"MasonPhan2110/Todo_Go_GPT/pkg/token"
	"MasonPhan2110/Todo_Go_GPT/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Config     utils.Config
	Store      db.Store
	TokenMaker token.Maker
	Router     *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		Config:     config,
		Store:      store,
		TokenMaker: tokenMaker,
	}

	r := gin.Default()

	server.AddRoutes(r)

	server.Router = r
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func (server *Server) AddRoutes(c *gin.Engine) {
	v1 := c.Group("api/v1")

	// v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	swagger := v1.Group("swagger")
	{
		swagger.GET("*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	auth := v1.Group("auth")
	{
		auth.POST("login", server.Login)
		auth.POST("renew_access", server.RenewAccessToken)
	}

	user := v1.Group("user")
	{
		user.POST("create", server.CreateUser)
		user.POST("update", server.UpdateUser)
	}

	todo := v1.Group("todo")
	{
		todo.POST("create", server.CreateTask)
		todo.GET("delete", server.DeleteTask)
		todo.GET("get", server.GetTask)
		todo.GET("list", server.ListTasks)
		todo.POST("update", server.UpdateTask)
	}

	// r.Use(middleware.AuthMiddleware(setting.AppSetting.TokenMaker))
}
