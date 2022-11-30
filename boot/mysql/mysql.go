package mysql

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"rustdesk-api-server/global"
)

// 注册Mysql驱动
func init() {
	orm.Debug = false
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Error("mysql 注册驱动失败:", err)
	}

	// 格式化连接符
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		global.ConfigVar.Mysql.Username,
		global.ConfigVar.Mysql.Password,
		global.ConfigVar.Mysql.Host,
		global.ConfigVar.Mysql.Port,
		global.ConfigVar.Mysql.Database,
	)

	// 注册链接数据库
	err = orm.RegisterDataBase("default", "mysql", connStr)
	if err != nil {
		logs.Error("mysql数据库注册失败", err)
	}

}
