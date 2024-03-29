package services

import (
	"github.com/beego/beego/v2/client/orm"
	"log"
	"rustdesk-api-server/app/models"
	"time"
)

var Token = new(TokenService)

type TokenService struct {
}

// 记录登录状态
func (t *TokenService) Login(user *models.User, clientId, uuid, token2 string) bool {
	m := orm.NewOrm()
	md := models.Token{
		Username:    user.Username,
		Uid:         user.Id,
		ClientId:    clientId,
		Uuid:        uuid,
		AccessToken: token2,
		ActiveTime:  time.Now().Unix(),
		LoginTime:   time.Now().Unix(),
		ExpireTime:  time.Now().Unix() + 3600,
	}
	//update, err := m.InsertOrUpdate(&md, "uid,client_id,uuid")
	// `sqlite3` nonsupport InsertOrUpdate in beego
	// 此orm不支持sqlite3执行 InsertOrUpdate, 拆改两步完成
	oldMd := models.Token{Uid: user.Id, ClientId: clientId, Uuid: uuid}
	_ = m.Read(&oldMd, "uid", "client_id", "uuid")
	rowId := int64(0)
	var err error
	// 存在主键更新
	if oldMd.Id != 0 {
		md.Id = oldMd.Id
		rowId, err = m.Update(&md)
		if err != nil {
			return false
		}
	}
	//不存在主键插入
	rowId, err = m.Insert(&md)
	if err != nil {
		return false
	}

	if rowId > 0 {
		log.Println("tokenUpdate", rowId)
	}
	return true
}

// 判断是否在线
func (u *TokenService) FindToken(uid int32, clientId, uuid string) *models.Token {
	o := orm.NewOrm()
	info := &models.Token{}
	err := o.QueryTable(new(models.Token)).Filter("uid", uid).Filter("client_id", clientId).Filter("uuid", uuid).One(info)
	if err != nil {
		return nil
	}

	return info

}

func (t *TokenService) Save(u *models.Token) bool {
	o := orm.NewOrm()
	update, err := o.Update(u)
	if err != nil || update == 0 {
		return false
	}
	return true
}

func (t *TokenService) FindTokens(uid int32) *[]models.Token {
	o := orm.NewOrm()
	var info []models.Token
	_, err := o.QueryTable(new(models.Token)).Filter("uid", uid).All(&info)
	if err != nil {
		return nil
	}

	return &info
}
