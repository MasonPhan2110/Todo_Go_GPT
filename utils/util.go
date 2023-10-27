package utils

import (
	"MasonPhan2110/Todo_Go_GPT/pkg/setting"
)

func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
