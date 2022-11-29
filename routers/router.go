package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"rustdesk-api-server/app/controllers"
)

// 初始化路由服务
func init() {
	// 跨域解决方案
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))

	// 设定路由信息
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/api/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/api/ab", &controllers.AddressBookController{}, "post:Update")
	beego.Router("/api/ab/get", &controllers.AddressBookController{}, "post:List")
	beego.Router("/api/audit", &controllers.AuditController{}, "post:Audit")
	beego.Router("/api/logout", &controllers.LogoutController{}, "post:Logout")
	beego.Router("/api/currentUser", &controllers.UserController{}, "post:CurrentUser")
	beego.Router("/api/reg", &controllers.UserController{}, "get:Reg")
	beego.Router("/api/set-pwd", &controllers.UserController{}, "get:SetPwd")

	// 设置错误路由
	beego.ErrorController(&controllers.ErrorController{})
}
