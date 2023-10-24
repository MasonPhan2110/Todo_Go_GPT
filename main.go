package main

import (
	db "MasonPhan2110/Todo_Go_GPT/db/setup"
	"MasonPhan2110/Todo_Go_GPT/docs"
	"MasonPhan2110/Todo_Go_GPT/pkg/setting"
	"MasonPhan2110/Todo_Go_GPT/routes"
	"MasonPhan2110/Todo_Go_GPT/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup("conf/app.ini")
	utils.Setup()

	dbSql := new(db.SqlDB)
	db.Setup(dbSql)
}

func main() {
	r := gin.Default()

	routes.AddRoutes(r)

	docs.SwaggerInfo.Title = "Swagger Todo List"
	docs.SwaggerInfo.Description = "Todo List Swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "0.0.0.0:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	if err := r.Run("0.0.0.0:8080"); err != nil {
		fmt.Println("Error in router run: ", err)
	}
}
