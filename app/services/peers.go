package services

import (
	"github.com/beego/beego/v2/client/orm"
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/models"
	"strings"
)

var Peers = new(PeersService)

type PeersService struct {
}

// 批量插入
func (this *PeersService) BatchAdd(uid int32, peers []dto.AbGetPeer) bool {
	if len(peers) == 0 {
		return false
	}

	tagList := []models.Peers{}
	for _, t := range peers {
		if len(t.Id) > 16 || t.Username == "----" {
			continue
		}

		tagList = append(tagList, models.Peers{
			Uid:      uid,
			ClientId: string(t.Id),
			Username: t.Username,
			Hostname: t.Hostname,
			Alias:    t.Alias,
			Platform: t.Platform,
			Tags:     strings.Join(t.Tags, ","),
		})
	}

	if len(tagList) == 0 {
		return true
	}

	_, err := orm.NewOrm().InsertMulti(3, tagList)
	if err != nil {
		return false
	}
	return true
}

// 查询用户名下Peers
func (this PeersService) FindPeers(uid int32) []models.Peers {
	ret := []models.Peers{}
	_, err := orm.NewOrm().QueryTable(new(models.Peers)).Filter("uid", uid).All(&ret)
	if err != nil {
		return nil
	}
	return ret
}

func (this PeersService) DeleteAll(uid int32) bool {
	_, err := orm.NewOrm().Raw("delete from rustdesk_peers where uid = ?", uid).Exec()
	if err != nil {
		return false
	}
	return true
}
