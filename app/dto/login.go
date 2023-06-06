package dto

//	{
//	   "username": "1g",
//	   "password": "21g2",
//	   "id": "1089363550",
//	   "uuid": "M0Y4MkI3N0MtMDMwMy01N0EwLTg5MzAtNDcwNUI4NUNFNUZD",
//	   "autoLogin": true,
//	   "type": "account",
//	   "verificationCode": "",
//	   "deviceInfo": {
//	       "os": "macos",
//	       "type": "client",
//	       "name": "xiaoyi510deimac.local"
//	   }
//	}
type LoginReq struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	Id               string `json:"id"`
	Uuid             string `json:"uuid"`
	AutoLogin        bool   `json:"autoLogin"`
	Type             string `json:"type"`
	VerificationCode string `json:"verificationCode"`
	DeviceInfo       struct {
		Os   string `json:"os"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"deviceInfo"`
}
