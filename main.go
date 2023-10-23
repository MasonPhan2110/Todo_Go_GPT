package main

import (
	db "MasonPhan2110/Todo_Go_GPT/db/setup"
	"MasonPhan2110/Todo_Go_GPT/pkg/setting"
	"MasonPhan2110/Todo_Go_GPT/utils"
)

func init() {
	setting.Setup("conf/app.ini")
	utils.Setup()

	dbSql := new(db.SqlDB)
	db.Setup(dbSql)
}

func main() {

}
