package routes

import (
	"MasonPhan2110/Todo_Go_GPT/api"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddRoutes(c *gin.Engine) {
	v1 := c.Group("api/v1")

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := v1.Group("user")
	{
		user.POST("login", api.Login)
		user.POST("create", api.CreateUser)
	}

	// r.Use(middleware.AuthMiddleware(setting.AppSetting.TokenMaker))
}
