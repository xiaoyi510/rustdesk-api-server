package controllers

import (
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/global"
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
	if req.AuthKey != global.ConfigVar.App.AuthKey {
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
	if req.AuthKey != global.ConfigVar.App.AuthKey {
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

// 分组
func (ctl *UserController) Users() {
	ctl.JSON(beegoHelper.H{
		"msg":   "成功",
		"total": 1,
		"data": []beegoHelper.H{
			{
				"name":     "默认用户",
				"email":    "ff",
				"note":     "哈哈",
				"status":   1,
				"is_admin": true,
			},
		},
	})
}

func (ctl *UserController) Peers() {
	ctl.JSON(beegoHelper.H{
		"msg":   "成功",
		"total": 1,
		"data": []beegoHelper.H{
			{
				"id": "test",
				"info": beegoHelper.H{
					"username": "",
					"os":       "", // windows
					//linux
					//macos
					//android
					"device_name": "",
				},
				"user":      "ff",
				"user_name": "占位",
				"node":      "tt",
				"is_admin":  true,
			},
		},
	})
}
