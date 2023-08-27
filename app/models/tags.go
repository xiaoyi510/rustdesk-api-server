package models

import (
	"github.com/beego/beego/v2/client/orm"
	"log"
)

type Tags struct {
	Id  int32  `json:"id"`
	Uid int32  `json:"uid"`
	Tag string `json:"tag"`
	Color string  `json:"color,omitempty"`
}

func (u *Tags) TableName() string {
	return "rustdesk_tags"
}

func init() {
	log.Printf("初始化模型")
	// 初始化模型
	orm.RegisterModel(new(Tags))
}
