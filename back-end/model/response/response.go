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

// 分类文件返回类型
type CategoryRes struct {
	FileID      uint      `json:"fileID"`
	FileName    string    `json:"fileName"`
	FileSize    int64     `json:"fileSize"`
	UpdateTime  time.Time `json:"updateTime"`
	DownloadNum int64     `json:"downloadNum"`
	FilePath    string    `json:"filePath"`
}

// 总的文件和文件夹返回类型
type FFRes struct {
	FileID      uint      `json:"fileID"`
	FileName    string    `json:"fileName"`
	FileSize    int64     `json:"fileSize"`
	UpdateTime  time.Time `json:"updateTime"`
	Type        string    `json:"type"`        // 是文件还是文件夹
	FileType    int       `json:"fileType"`    // 哪种文件 0:文件夹 1:图片 2:视频 3:音频 4:文档 5:其他
	DownloadNum int64     `json:"downloadNum"` // 如果是文件夹就忽略此字段
	FilePath    string    `json:"filePath"`
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
var ErrorAvatarUpload Error = Error{"头像上传失败"}
var ErrorFileUpload Error = Error{"文件上传失败"}
var ErrorFilePathNotExist Error = Error{"文件路径不存在"}
var ErrorFileExist Error = Error{"同名文件已存在"}
var ErrorStoreFull Error = Error{"网盘已满，无法上传"}

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
