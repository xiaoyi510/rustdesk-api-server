package controllers

import (
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/utils/beegoHelper"
	"time"
)

var Audit = new(AuditController)

type AuditController struct {
	BaseController
}

// 心跳检测 POST
func (ctl *AuditController) Audit() {
	req := dto.AuditReq{}

	if err := ctl.BindJSON(&req); err != nil {
		ctl.JSON(beegoHelper.H{
			"error": "请求参数异常",
		})
		return
	}

	// 设置当前用户 在线信息
	tokenInfo := services.Token.FindToken(req.Id, req.Id1, req.Uuid)
	if tokenInfo != nil {
		// 修改token活跃时间
		tokenInfo.ActiveTime = time.Now().Unix()
		if !services.Token.Save(tokenInfo) {
			ctl.JSON(beegoHelper.H{
				"error": "保存登录心跳错误",
			})
		}

		ctl.JSON(beegoHelper.H{
			"data": "在线",
		})
	} else {
		ctl.JSON(beegoHelper.H{
			"error": "设备异常",
		})
	}

}
