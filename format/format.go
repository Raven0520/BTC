package format

import (
	"github.com/raven0520/btc/util"
)

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
