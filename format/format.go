package format

import (
	"github.com/raven0520/btc/util"
	"github.com/raven0520/proto/btc"
)

// Response 简单返回
type Response struct {
	Message string `json:"message"`
	Data    interface{}
}

// BoolResponse 返回结果
type BoolResponse struct {
	Message string `json:"message"`
	Data    bool   `json:"data"`
}

// ResultResponse Return true / false result
func ResultResponse(result bool, msg string, err error) (*btc.BoolResponse, error) {
	return &btc.BoolResponse{
		Message: Message(msg, err),
		Data:    result,
	}, nil
}

// RequestOK 通用响应成功返回
func RequestOK(err error) string {
	return Message(util.RequestOK, err)
}

// Message 通用返回响应消息
func Message(message string, err error) string {
	if err != nil {
		return err.Error()
	}
	return message
}
