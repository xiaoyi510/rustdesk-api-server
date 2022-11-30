package services

import (
	"github.com/beego/beego/v2/client/orm"
	"rustdesk-api-server/app/models"
	"rustdesk-api-server/global"
	"rustdesk-api-server/utils/gmd5"
	"time"
)

var User = new(UserService)

type UserService struct {
}

func (u *UserService) Reg(username, password string) bool {
	// 生成密码
	hashPwd := u.GenPwd(password)

	// 插入或者修改
	m := &models.User{
		Username:   username,
		Password:   hashPwd,
		Status:     1,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	_, err := orm.NewOrm().Insert(m)
	if err != nil {
		return false
	}

	return true
}

// 生成保存的密码
func (u *UserService) GenPwd(password string) string {
	// 校验密码是否正确
	pwd, err := gmd5.Encrypt(password + global.ConfigVar.App.CryptKey)
	if err != nil {
		panic("md5 encrypt Err" + err.Error())
	}

	return pwd
}

// 重置密码
func (u *UserService) ResetPassword(username string, password string) bool {
	// 生成密码
	hashPwd := u.GenPwd(password)

	m := User.FindByUserName(username)
	if m == nil {
		return false
	}
	m.Password = hashPwd
	_, err := orm.NewOrm().Update(m, "password")
	if err != nil {
		return false
	}

	return true
}

// 根据用户名查询用户信息
func (u *UserService) FindByUserName(username string) *models.User {
	ret := models.User{}
	err := orm.NewOrm().QueryTable(new(models.User)).Filter("username", username).One(&ret)
	if err != nil {
		return nil
	}
	return &ret
}

func (u *UserService) Logout(info *models.User, clientId string) bool {
	// 删除登录的token
	token := &models.Token{}

	_, err := orm.NewOrm().Raw("delete from "+token.TableName()+" where uid = ? and client_id = ?", info.Id, clientId).Exec()
	if err != nil {
		return false
	}
	return true
}
