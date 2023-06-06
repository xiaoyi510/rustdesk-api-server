package controllers

import (
	"encoding/json"
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/utils/beegoHelper"
	"strings"
)

var Address = new(AddressBookController)

type AddressBookController struct {
	BaseController
}

// 查看地址谱列表
func (ctl *AddressBookController) List() {
	ack := dto.AbGetAck{}
	ack.Tags = []string{}
	// 查询 tags
	tags := services.Tags.FindTags(ctl.loginUserInfo.Id)
	for _, item := range tags {
		ack.Tags = append(ack.Tags, item.Tag)
	}

	// 查询 peers
	ack.Peers = []dto.AbGetPeer{}
	peerDbs := services.Peers.FindPeers(ctl.loginUserInfo.Id)
	for _, item := range peerDbs {
		ack.Peers = append(ack.Peers, dto.AbGetPeer{
			Id:       item.ClientId,
			Username: item.Username,
			Hostname: item.Hostname,
			Alias:    item.Alias,
			Platform: item.Platform,
			Tags:     strings.Split(item.Tags, ","),
		})
	}

	// 查询出来所有已登录的账号列表
	tokens := services.Token.FindTokens(ctl.loginUserInfo.Id)
	for _, item := range *tokens {
		ist := false
		for _, bookItem := range ack.Peers {
			if bookItem.Id == item.ClientId {
				ist = true
				break
			}
		}
		if !ist {
			ack.Peers = append(ack.Peers, dto.AbGetPeer{
				Id:       item.ClientId,
				Username: "----",
				Hostname: item.ClientId,
				Alias:    "本号归属:" + item.ClientId,
				Platform: "无",
				Tags:     strings.Split("", ","),
			})
		}
	}

	jdata, _ := json.Marshal(ack)

	ctl.JSON(beegoHelper.H{
		//"error":     false,
		"data": string(jdata),
		//"update_at": time.Now().Format("2006-01-02 15:04:05"),
	})
}

// 更新地址谱
func (ctl *AddressBookController) Update() {
	req := dto.AbUpdateReq{}

	if err := ctl.BindJSON(&req); err != nil {
		ctl.JSON(beegoHelper.H{
			"error": "请求参数异常",
		})
		return
	}

	// 解析数据
	reqSub := &dto.AbUpdateSub{}
	err := json.Unmarshal([]byte(req.Data), reqSub)
	if err != nil {
		ctl.JSON(beegoHelper.H{
			"error": "请求数据异常",
		})
	}

	// 批量删除tags
	services.Tags.DeleteAll(ctl.loginUserInfo.Id)
	// 批量删除Peers
	services.Peers.DeleteAll(ctl.loginUserInfo.Id)

	// 开始批量插入tags
	if !services.Tags.BatchAdd(ctl.loginUserInfo.Id, reqSub.Tags) {
		ctl.JSON(beegoHelper.H{
			"error": "导入标签失败",
		})
	}
	// 开始批量插入peers
	if !services.Peers.BatchAdd(ctl.loginUserInfo.Id, reqSub.Peers) {
		ctl.JSON(beegoHelper.H{
			"error": "导入地址簿失败",
		})
	}

	ctl.JSON(beegoHelper.H{
		"data": "成功",
	})

}
