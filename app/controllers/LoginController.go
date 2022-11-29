package controllers

import (
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/utils/beegoHelper"
	"rustdesk-api-server/utils/common"
	"strings"
)

var Login = new(LoginController)

type LoginController struct {
	BaseController
}

// 登录
func (ctl *LoginController) Login() {
	if ctl.Ctx.Input.IsPost() {
		// 获取请求参数
		var req dto.LoginReq
		if err := ctl.BindJSON(&req); err != nil {
			ctl.JSON(common.JsonResult{
				Error: err.Error(),
			})
		}
		req.Username = strings.TrimSpace(req.Username)
		if len(req.Username) == 0 {
			ctl.JSON(common.JsonResult{
				Code:  -1,
				Error: "用户名不能为空",
			})
		}
		req.Password = strings.TrimSpace(req.Password)
		if len(req.Password) == 0 {
			ctl.JSON(common.JsonResult{
				Code:  -1,
				Error: "密码不能为空",
			})
		}
		req.ClientId = strings.TrimSpace(req.ClientId)
		if len(req.ClientId) == 0 {
			ctl.JSON(common.JsonResult{
				Code:  -1,
				Error: "客户端ID不能为空",
			})
		}

		// 查询数据库中的账号密码是否合法
		token, err := services.Login.UserLogin(req.Username, req.Password, req.ClientId, req.Uuid, ctl.Ctx)
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code:  -1,
				Error: err.Error(),
			})
		}

		ctl.JSON(beegoHelper.H{
			"access_token": token,
			"user": beegoHelper.H{
				"name": req.Username,
			},
		})

	}

}
