package dto

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ClientId string `json:"id"`
	Uuid     string `json:"uuid"`
}
