package models

import (
	"github.com/beego/beego/v2/client/orm"
	"log"
)

type Peers struct {
	DeviceId int32  `json:"device_id" orm:"column(deviceid);auto"`
	Uid      int32  `json:"uid"`
	ClientId string `json:"id" orm:"column(id)"`
	Username string `json:"username"`
	Hostname string `json:"hostname"`
	Alias    string `json:"alias"`
	Platform string `json:"platform"`
	Tags     string `json:"tags"`
}

func (u *Peers) TableName() string {
	return "rustdesk_peers"
}

func init() {
	log.Printf("初始化模型")
	// 初始化模型
	orm.RegisterModel(new(Peers))
}
