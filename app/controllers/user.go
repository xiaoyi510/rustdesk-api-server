package controllers

import (
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/configer"
	"rustdesk-api-server/utils/beegoHelper"
)

type UserController struct {
	BaseController
}

// 当前用户信息
func (ctl *UserController) CurrentUser() {
	ctl.JSON(beegoHelper.H{
		"name": ctl.loginUserInfo.Username,
	})
}

// 注册用户
func (ctl *UserController) Reg() {
	req := dto.UserRegReq{}
	req.Username = ctl.GetString("username")
	req.Password = ctl.GetString("password")
	req.AuthKey = ctl.GetString("auth_key")
	if len(req.Username) < 4 || len(req.Username) > 20 {
		ctl.JSON(beegoHelper.H{
			"error": "用户名在4-20位之间",
		})
	}

	if len(req.AuthKey) == 0 {
		ctl.JSON(beegoHelper.H{
			"error": "请输入授权码",
		})
	}

	// 判断注册密钥是否合法
	if req.AuthKey != configer.ConfigVar.App.AuthKey {
		ctl.JSON(beegoHelper.H{
			"error": "授权码错误",
		})
	}

	// 去注册账号
	if services.User.Reg(req.Username, req.Password) {
		ctl.JSON(beegoHelper.H{
			"msg": "注册成功",
		})
	} else {
		ctl.JSON(beegoHelper.H{
			"error": "注册失败",
		})
	}
}

// 修改用户密码
func (ctl *UserController) SetPwd() {
	req := dto.UserSetPwdReq{}
	req.Username = ctl.GetString("username")
	req.Password = ctl.GetString("password")
	req.AuthKey = ctl.GetString("auth_key")
	if len(req.Username) < 4 || len(req.Username) > 20 {
		ctl.JSON(beegoHelper.H{
			"error": "用户名在4-20位之间",
		})
	}

	if len(req.AuthKey) == 0 {
		ctl.JSON(beegoHelper.H{
			"error": "请输入授权码",
		})
	}

	// 判断注册密钥是否合法
	if req.AuthKey != configer.ConfigVar.App.AuthKey {
		ctl.JSON(beegoHelper.H{
			"error": "授权码错误",
		})
	}

	// 去注册账号
	if services.User.ResetPassword(req.Username, req.Password) {
		ctl.JSON(beegoHelper.H{
			"msg": "修改成功",
		})
	} else {
		ctl.JSON(beegoHelper.H{
			"error": "修改失败",
		})
	}
}
