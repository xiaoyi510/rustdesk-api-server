package controllers

import "rustdesk-api-server/utils/beegoHelper"

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.JSON(beegoHelper.H{
		"error": "未找到页面",
	})
	//c.Ctx.WriteString("customize page not found")
	//c.Data= "page not found"
	//c.TplName = "404.tpl"
}

func (c *ErrorController) Error501() {
	c.JSON(beegoHelper.H{
		"error": "服务端报错",
	})
	//c.Ctx.WriteString("customize server error")
	//c.Data["content"] = "server error"
	//c.TplName = "501.tpl"
}
