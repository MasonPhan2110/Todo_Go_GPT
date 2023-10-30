package setup

import (
	sqlc "MasonPhan2110/Todo_Go_GPT/db/sqlc"
	"database/sql"
	"fmt"
)

type setupDB interface {
	Init(string, string) error
}

type SqlDB struct {
	db *sql.DB
}

func (s *SqlDB) Init(DB_DRIVER, DB_SOURCE string) error {
	var err error

	s.db, err = sql.Open(DB_DRIVER, DB_SOURCE)
	fmt.Println("DB Driver: ", DB_DRIVER)
	fmt.Println("DB Source: ", DB_SOURCE)
	if err != nil {
		return err
	}

	sqlc.DBStore = sqlc.NewStore(s.db)

	return err
}

func Setup(dbConn setupDB, DB_DRIVER, DB_SOURCE string) {
	if err := dbConn.Init(DB_DRIVER, DB_SOURCE); err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Connected")
}
