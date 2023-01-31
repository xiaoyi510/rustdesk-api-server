package global

import "rustdesk-api-server/global/confDto"

type Config struct {
	DBType string            `json:"dbtype"`
	Mysql  confDto.Mysql     `json:"mysql"`
	App    confDto.AppConfig `json:"app"`
}
