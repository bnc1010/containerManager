package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ResponseOK 成功响应体
func ResponseOK(c *app.RequestContext, msg string, data interface{}) {
	c.JSON(consts.StatusOK, utils.H{
		"code":    10200,
		"message": msg,
		"data":    data,
	})
}

// ResponseError 错误响应体
func ResponseError(c *app.RequestContext, msg string, err error) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    10500,
		"message": msg,
		"data": utils.H{
			"err_msg": errMsg,
		},
	})
}

// ResponseForbid 禁止响应体
func ResponseForbid(c *app.RequestContext, msg string) {
	c.JSON(consts.StatusOK, utils.H{
		"code":   10403,
		"message": msg,
	})
}

// ResponseNotFound 404响应体
func ResponseNotFound(c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"code":   10404,
		"message": "Not Found",
	})
}