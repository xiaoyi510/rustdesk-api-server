package configer

import "rustdesk-api-server/configer/config"

type Config struct {
	Mysql config.Mysql     `json:"mysql"`
	App   config.AppConfig `json:"app"`
}
