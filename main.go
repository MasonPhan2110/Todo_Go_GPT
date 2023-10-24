package main

import (
	db "MasonPhan2110/Todo_Go_GPT/db/setup"
	"MasonPhan2110/Todo_Go_GPT/middleware"
	"MasonPhan2110/Todo_Go_GPT/pkg/setting"
	"MasonPhan2110/Todo_Go_GPT/routes"
	"MasonPhan2110/Todo_Go_GPT/utils"

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
	r.Use(middleware.AuthMiddleware(setting.AppSetting.TokenMaker))
	routes.AddRoutes(r)
}
