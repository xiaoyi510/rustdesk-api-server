package common

type JsonResult struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count,omitempty"`
}
