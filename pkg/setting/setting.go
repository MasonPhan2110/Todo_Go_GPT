package setting

import (
	"MasonPhan2110/Todo_Go_GPT/pkg/token"
	"fmt"
	"log"
	"time"

	"gopkg.in/ini.v1"
)

type App struct {
	Environment          string
	JwtSecret            string
	TokenSymmetricKey    string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	TokenMaker           token.Maker
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

var PostgreDBSettings = &PostgresDB{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup(path string) {
	var err error
	cfg, err = ini.Load(path)

	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	tokenMaker, err := token.NewPasetoMaker(AppSetting.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("setting.Setup, fail to create token maker': %v", err)
	}

	AppSetting.TokenMaker = tokenMaker

	mapTo("PostgreDatabase", PostgreDBSettings)
	// DB_SOURCE=postgresql://root:rootroot@localhost:5432/nft_poc?sslmode=disable

	// PostgreSQL db source
	PostgreDBSettings.DBSource = fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable",
		PostgreDBSettings.Type, PostgreDBSettings.User, PostgreDBSettings.Password, PostgreDBSettings.Host, PostgreDBSettings.Port, PostgreDBSettings.DBName)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
