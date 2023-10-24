package routes

import (
	"MasonPhan2110/Todo_Go_GPT/api"

	"github.com/gin-gonic/gin"
)

func AddRoutes(c *gin.Engine) {
	v1 := c.Group("api/v1")

	user := v1.Group("user")
	{
		user.POST("login", api.Login)
	}
}
