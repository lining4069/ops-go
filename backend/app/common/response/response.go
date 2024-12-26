package response

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/ops-go/backend/global"
	"net/http"
)

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 响应成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "ok",
		Data:    data,
	})
}

// Fail 失败响应 code不为0 表示失败
func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    errorCode,
		Message: msg,
		Data:    nil,
	})
}

// FailByError 失败响应 返回自定义错误的错误码，错误信息
func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}