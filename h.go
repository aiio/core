package core

import (
	"github.com/aiio/core/codex"
	"github.com/aiio/core/valid"

	"github.com/go-playground/validator/v10"
)

// H common struct
type H struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Other interface{} `json:"other,omitempty"`
}

// Resp Resp
func Resp(code int, msg string, data interface{}) *H {
	return &H{Code: code, Msg: msg, Data: data}
}

func AuthErr(msg string) *H {
	return &H{Code: codex.AuthErr, Msg: msg}
}

func Response(code int, msg string, data, other interface{}) *H {
	return &H{Code: code, Msg: msg, Data: data, Other: other}
}

func List(data, other interface{}) *H {
	return &H{Code: codex.Success, Msg: "success", Data: data, Other: other}
}

func Success() *H {
	return &H{Code: codex.Success, Msg: "success"}
}

// Data Data
func Data(data interface{}) *H {
	return &H{Code: codex.Success, Msg: "success", Data: data}
}

// NewH NewH
func NewH(data interface{}) *H {
	return &H{Code: codex.Success, Msg: "success", Data: data}
}

// NewE NewE
func NewE(err error) *H {
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		return &H{Code: codex.ValidErr, Msg: "参数错误,详情请看data内容", Data: errs.Translate(valid.Trans)}
	}
	return &H{Code: codex.Error, Msg: err.Error()}
}
