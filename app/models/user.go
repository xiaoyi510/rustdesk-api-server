package models

import (
	"github.com/beego/beego/v2/client/orm"
	"log"
)

type User struct {
	Id            int32  `json:"id"`
	Username      string `json:"username" orm:"unique"`
	Password      string `json:"password"`
	Status        int32  `json:"status"`
	LastLoginTime int64  `json:"last_login_time"`
	LastLoginIp   string `json:"last_login_ip"`
	CreateTime    int64  `json:"create_time"`
	UpdateTime    int64  `json:"update_time"`
	DeleteTime    int64  `json:"delete_time,omitempty"`
}

func (u *User) TableName() string {
	return "rustdesk_users"
}

func init() {
	log.Printf("初始化模型")
	// 初始化模型
	orm.RegisterModel(new(User))
}
