package setup

import (
	sqlc "MasonPhan2110/Todo_Go_GPT/db/sqlc"
	"MasonPhan2110/Todo_Go_GPT/pkg/setting"
	"database/sql"
	"fmt"
)

type setupDB interface {
	Init() error
}

type SqlDB struct {
	db *sql.DB
}

func (s *SqlDB) Init() error {
	var err error

	s.db, err = sql.Open(setting.PostgresDBSetting.DBDriver, setting.PostgresDBSetting.DBSource)
	fmt.Println("DB Driver: ", setting.PostgresDBSetting.DBDriver)
	fmt.Println("DB Source: ", setting.PostgresDBSetting.DBSource)
	if err != nil {
		return err
	}

	sqlc.DBStore = sqlc.NewStore(s.db)

	return err
}

func Setup(dbConn setupDB) {
	if err := dbConn.Init(); err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Connected")
}
