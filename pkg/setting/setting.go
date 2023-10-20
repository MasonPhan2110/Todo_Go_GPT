package setting

import (
	"MasonPhan2110/Todo_Go_GPT/pkg/token"
	"time"
)

type App struct {
	Environment         string
	JwtSecret           string
	TokenSymmetricKey   string
	AccessTokenDuration time.Duration
	TokenMaker          token.Maker
}

type PostgresDB struct {
	Type     string
	DBDriver string
	DBSource string
	DBName   string
	User     string
	Password string
	Host     string
	Port     string
}

var AppSetting = &App{}
var PostgresDBSetting = &PostgresDB{}

// Setup initialize the configuration instance
func Setup(path string) {}
