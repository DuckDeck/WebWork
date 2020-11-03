package model

type Result struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	CMsg     string      `json:"cmsg"`
	MsgDebug string      `json:"MsgDebug"` //用于debug模式
	Count    int         `json:"count"`
	Data     interface{} `json:"data"`
}

//错误码
//负数请求过程发生的错误
//正数表示业务错误
//表示请求参数错了-100
