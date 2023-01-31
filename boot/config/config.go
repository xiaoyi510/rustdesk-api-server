package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"rustdesk-api-server/constant"
	"rustdesk-api-server/global"
)

// 加载配置项
func init() {
	fmt.Println("加载配置项")
	var config string
	if configEnv := os.Getenv(constant.ConfigEnv); configEnv == "" {
		config = constant.ConfigFile
	} else {
		config = configEnv
	}

	// 判断是否有配置文件
	_, err := os.Stat(config)
	if err != nil && os.IsNotExist(err) {
		// 配置文件不存在
		err := os.WriteFile(config, []byte(`dbtype: 'sqlite3'
mysql:
  host: '127.0.0.1'
  port: 3306
  database: 'rustdesk'
  username: 'root'
  password: ''
app:
  authkey: 123456
  cryptkey: NanEVhjEwuPSemoJkwcYEcjDJRVWcZfb9bIIZcBeswhPP`), 0777)
		if err != nil {
			panic(err)
		}
	}

	v := viper.New()

	// 设置配置文件
	v.SetConfigFile(config)

	// 读取配置
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("加载配置文件失败: %s", err))
	}

	// 监控配置更新
	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改:", in.Name)
		if err := v.Unmarshal(&global.ConfigVar); err != nil {
			panic(err)
		}
	})

	if err := v.Unmarshal(&global.ConfigVar); err != nil {
		panic(err)
	}
	fmt.Println("加载配置项 完成")

}
