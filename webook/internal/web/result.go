package web

// 501001=>代表验证码没有发送
type Result struct {
	//这个是业务错误码
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
