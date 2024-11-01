// 定义各种数据结构
package request

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Code     string `json:"code" form:"code"` // 验证码
}

type VerifyReq struct {
	Email string `json:"email"`
}

type UploadReq struct {
	FilePath string `json:"filepath" form:"filepath"`
}

type CreateFolderReq struct {
	FolderName string `json:"folderName"`
}
