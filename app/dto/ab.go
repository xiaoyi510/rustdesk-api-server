package dto

type AbTag_colors map[string]int64

type AbGetAck struct {
	Tags  []string    `json:"tags"`
	Peers []AbGetPeer `json:"peers"`
	Tag_colors string `json:"tag_colors,omitempty"`	
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
	Tag_colors string `json:"tag_colors,omitempty"`	
}
