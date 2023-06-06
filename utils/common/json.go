package common

type JsonResult struct {
	Code  int         `json:"code,omitempty"`
	Msg   string      `json:"msg,omitempty"`
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Count int64       `json:"count,omitempty"`
}
