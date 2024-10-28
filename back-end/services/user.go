// 服务层
package services

import (
	"disk/database"
	"disk/model"
	"disk/model/request"
	"disk/model/response"
	"disk/utils"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// 登录
func Login(ctx *gin.Context) {
	var loginReq request.LoginReq
	// 参数绑定失败
	if err := ctx.BindJSON(&loginReq); err != nil {
		response.RequestError(ctx, response.ErrorParams)
		return
	}
	// 检验用户是否存在
	var user *model.User
	if user = database.DBFindByUsername(loginReq.Username); user == nil {
		response.RequestError(ctx, response.ErrorUserNotFound)
		return
	}
	// 检验密码是否相同
	if err := utils.CheckPassword(loginReq.Password, user.Password); err != nil {
		response.RequestError(ctx, response.ErrorUserNotFound)
		return
	}
	// 生成token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		response.ServerError(ctx, response.ErrorTokenGenerate)
		return
	}
	// 成功返回
	response.Success(ctx, "登录成功", gin.H{
		"avatar":   user.Avatar,
		"username": user.Username,
		"email":    user.Email,
		"token":    token,
	})
}

// 注册
func SignUp(ctx *gin.Context) {
	var signUpReq request.SignUpReq
	// 必须得用shouldbind, 因为需要既绑定JSON，又要绑定formData
	if err := ctx.ShouldBind(&signUpReq); err != nil {
		response.RequestError(ctx, response.ErrorParams)
		return
	}
	fmt.Println(signUpReq)
	if signUpReq.Username == "" || signUpReq.Password == "" || signUpReq.Email == "" || signUpReq.Code == "" {
		response.RequestError(ctx, response.ErrorParams)
		return
	}
	// 检验验证码是否正确
	if code, err := database.Get(signUpReq.Email); err != nil {
		fmt.Println("get redis code failed")
		response.ServerError(ctx, response.ErrorCodeExpired)
		return
	} else if code == "" {
		response.ServerError(ctx, response.ErrorCodeExpired)
		return
	} else if code != signUpReq.Code {
		response.RequestError(ctx, response.ErrorCodeIncorrect)
		return
	}
	// 检验用户名和邮箱是否存在
	if err := database.CheckUsernameAndEmail(signUpReq.Username, signUpReq.Email); err != nil {
		fmt.Println("username or email exist")
		response.RequestError(ctx, *err)
		return
	}
	// 获取头像
	var avatar string
	file, err := ctx.FormFile("avatar")
	if err != nil {
		fmt.Println("use default avatar")
		// 使用默认头像
		avatar = utils.GetImageUrl("default-avatar.png")
	} else {
		dir, _ := os.Getwd()
		filename := file.Filename
		filepath := dir + "/static/images/" + filename
		if err := ctx.SaveUploadedFile(file, filepath); err != nil {
			fmt.Println("avatar save failed")
		} else {
			avatar = utils.GetImageUrl(filename)
		}
	}
	// 加密密码
	hashedPassword, err := utils.EncryptPassword(signUpReq.Password)
	if err != nil {
		response.ServerError(ctx, response.ErrorPasswordEncrypt)
		return
	}
	user := &model.User{
		Avatar:   avatar,
		Username: signUpReq.Username,
		Password: hashedPassword,
		Email:    signUpReq.Email,
	}
	if err := database.DBCreateUserAndFileStore(user); err != nil {
		// 用户创建有问题，就需要把保存在本地的头像文件删除掉，避免浪费资源
		if avatar != "" {
			dir, _ := os.Getwd()
			filename := file.Filename
			filepath := dir + "/static/images" + filename
			os.Remove(filepath)
		}
		response.ServerError(ctx, response.ErrorUserCreate)
		return
	}
	response.Success(ctx, "注册成功", nil)
}

// 发送验证码验证
func Verify(ctx *gin.Context) {
	var verifyReq request.VerifyReq
	if err := ctx.BindJSON(&verifyReq); err != nil {
		response.RequestError(ctx, response.ErrorParams)
		return
	}
	// 发送验证码
	if err := utils.SendEmail(verifyReq.Email); err != nil {
		response.ServerError(ctx, response.ErrorEmailSend)
		return
	}
	response.Success(ctx, "验证码发送成功", nil)
}

// 更换头像
func ChangeAvatar(ctx *gin.Context) {
	// 获取用户id
	userID, ok := ctx.Get("userID")
	if !ok {
		response.UnauthorizedError(ctx, response.ErrorUnauthorized)
	}
	// 获取头像
	file, err := ctx.FormFile("avatar")
	if err != nil {
		response.RequestError(ctx, response.ErrorAvatarUpload)
		return
	}
	// 创建头像本地地址映射url
	var avatar string
	dir, _ := os.Getwd()
	suffix := strings.Split(file.Filename, ".")
	avatarName := fmt.Sprintf("%s_%d_%s%s", "user", userID.(uint), "avatar", "."+suffix[len(suffix)-1])
	filepath := dir + "/static/images/" + avatarName
	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		response.ServerError(ctx, response.ErrorAvatarUpload)
		return
	} else {
		avatar = utils.GetImageUrl(avatarName)
	}
	// 把用户查找出来并更改对应的头像地址
	// step1: 找出用户
	// step2: 删除用户之前的头像
	// step3: 更新用户的头像
	// step4: 返回更新之后的头像url地址
	oldAvatar := database.UpdateAvatar(userID.(uint), avatar)
	// 删除头像
	oldAvatarPath := dir + "/static/images/" + oldAvatar
	os.Remove(oldAvatarPath)
	response.Success(ctx, "头像更新成功", gin.H{
		"avatar": avatar,
	})
}
