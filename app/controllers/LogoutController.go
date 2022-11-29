package controllers

import (
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/utils/beegoHelper"
)

type LogoutController struct {
	BaseController
}

// 退出登录
func (ctl *LogoutController) Logout() {
	req := dto.LogoutReq{}
	if err := ctl.BindJSON(&req); err != nil {
		ctl.JSON(beegoHelper.H{
			"error": err.Error(),
		})
	}

	if services.User.Logout(ctl.loginUserInfo, req.Id) {
		ctl.JSON(beegoHelper.H{"data": "退出完成"})
	} else {
		ctl.JSON(beegoHelper.H{"error": "退出失败"})
	}
}
