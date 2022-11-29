package dto

type AbGetAck struct {
	Tags  []string    `json:"tags"`
	Peers []AbGetPeer `json:"peers"`
}

type AbGetPeer struct {
	Id       string   `json:"id"`
	Username string   `json:"username"`
	Hostname string   `json:"hostname"`
	Alias    string   `json:"alias"`
	Platform string   `json:"platform"`
	Tags     []string `json:"tags"`
}

type AbUpdateReq struct {
	Data string `json:"data"`
}

type AbUpdateSub struct {
	Tags  []string    `json:"tags"`
	Peers []AbGetPeer `json:"peers"`
}
