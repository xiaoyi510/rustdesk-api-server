package models

import (
	"github.com/beego/beego/v2/client/orm"
	"log"
)

type Token struct {
	Id          int32  `json:"id" orm:"auto"`
	Username    string `json:"username"`
	Uid         int32  `json:"uid"`
	ClientId    string `json:"client_id"`
	Uuid        string `json:"uuid"`
	AccessToken string `json:"access_token"`
	ActiveTime  int64  `json:"login_time"`
	LoginTime   int64  `json:"active_time"`
	ExpireTime  int64  `json:"expire_time"`
}

func (u *Token) TableName() string {
	return "rustdesk_token"
}

// 多字段唯一键
func (u *Token) TableUnique() [][]string {
	return [][]string{
		[]string{"uid", "client_id", "uuid"},
	}
}

func init() {
	log.Printf("初始化模型")
	// 初始化模型
	orm.RegisterModel(new(Token))
}
