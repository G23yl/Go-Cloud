// 响应数据报格式
package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Error struct {
	msg string
}

type CategoryRes struct {
	FileID      uint      `json:"fileID"`
	FileName    string    `json:"fileName"`
	FileSize    int64     `json:"fileSize"`
	UpdateTime  time.Time `json:"updateTime"`
	DownloadNum int64     `json:"downloadNum"`
	FileUrl     string    `json:"fileUrl"`
}

var ErrorParams Error = Error{"参数绑定错误"}
var ErrorUserNotFound Error = Error{"用户名或密码错误"}
var ErrorTokenGenerate Error = Error{"token生成失败"}
var ErrorUnauthorized Error = Error{"token失效"}
var ErrorUsernameExist Error = Error{"用户名已存在"}
var ErrorEmailExist Error = Error{"每个邮箱只能注册一次"}
var ErrorPasswordEncrypt = Error{"密码加密失败"}
var ErrorUserCreate Error = Error{"用户创建失败"}
var ErrorEmailSend Error = Error{"邮件发送失败"}
var ErrorCodeExpired Error = Error{"验证码已过期"}
var ErrorCodeIncorrect Error = Error{"验证码错误"}

func newResponse(code int, msg string, data any) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

func Success(ctx *gin.Context, message string, data any) {
	ctx.JSON(http.StatusOK, newResponse(http.StatusOK, message, data))
}

func RequestError(ctx *gin.Context, err Error) {
	ctx.JSON(http.StatusBadRequest, newResponse(http.StatusBadRequest, err.msg, nil))
}

func ServerError(ctx *gin.Context, err Error) {
	ctx.JSON(http.StatusInternalServerError, newResponse(http.StatusInternalServerError, err.msg, nil))
}

func UnauthorizedError(ctx *gin.Context, err Error) {
	ctx.JSON(http.StatusUnauthorized, newResponse(http.StatusUnauthorized, err.msg, nil))
}
