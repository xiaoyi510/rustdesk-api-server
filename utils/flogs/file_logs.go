package flogs

import (
	"encoding/json"
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func InitLogger() {
	logConf := make(map[string]interface{})
	logConf["filename"], _ = beego.AppConfig.String("log_path")
	logConf["level"], _ = beego.AppConfig.Int("log_level")
	logConf["maxlines"], _ = beego.AppConfig.Int("maxlines")
	logConf["maxsize"], _ = beego.AppConfig.Int("maxsize")

	confStr, _ := json.Marshal(logConf)
	logs.SetLogger(logs.AdapterFile, string(confStr))
	logs.SetLogFuncCall(true)
	return
}
