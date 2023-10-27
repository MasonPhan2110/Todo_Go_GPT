package main

import (
	"MasonPhan2110/Todo_Go_GPT/db/setup"
	db "MasonPhan2110/Todo_Go_GPT/db/sqlc"
	"MasonPhan2110/Todo_Go_GPT/docs"
	"MasonPhan2110/Todo_Go_GPT/pkg/setting"
	"MasonPhan2110/Todo_Go_GPT/server"
	"MasonPhan2110/Todo_Go_GPT/utils"
	"fmt"
)

func init() {
	setting.Setup("conf/app.ini")
	utils.Setup()

	dbSql := new(setup.SqlDB)
	setup.Setup(dbSql)
}

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load config")
		return
	}

	dbSql := new(setup.SqlDB)
	setup.Setup(dbSql)

	docs.SwaggerInfo.Title = "Swagger Todo List"
	docs.SwaggerInfo.Description = "Todo List Swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "0.0.0.0:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	runGinServer(config, db.DBStore)
}

func runGinServer(config utils.Config, store db.Store) {
	server, err := server.NewServer(config, store)
	if err != nil {
		fmt.Println("Cannot create Server: ", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		fmt.Println("Cannot start server: ", err)
	}
}
