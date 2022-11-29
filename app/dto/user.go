package dto

// 用户注册
type UserRegReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AuthKey  string `json:"auth_key"`
}

// 用户修改密码
type UserSetPwdReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AuthKey  string `json:"auth_key"`
}
