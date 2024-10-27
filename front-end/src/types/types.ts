// 定义数据结构

export interface LoginForm {
  username: string
  password: string
}

// 基础响应包含code和msg，但是data字段根据具体情况返回对应类型
export interface BasicResponse {
  code: number
  msg: string
  data: any
}

// 继承BasicResponse，根据具体返回情况设置data字段的类型
export interface User {
  avatar: string
  email: string
  token: string
  username: string
}
export interface LoginResponse extends BasicResponse {
  data: User
}

// 菜单条目
export interface MenuItem {
  index: string
  itemName: string
}
export type MenuItems = MenuItem[]

// 注册表单
export interface SignUpInfo {
  username: string
  password: string
  confirmPassword: string
  code: string
  email: string
}

// 注册信息
export type SignUpForm = Omit<SignUpInfo, "confirmPassword">

// 注册表单检验规则
export type SignUpFormRules = Omit<SignUpInfo, "code">

// index页面响应
export interface OverallData {
  filesNum: number
  foldersNum: number
  storeDetail: {
    capacity: number
    size: number
  }
  categoryNum: {
    DocNum: number
    ImgNum: number
    VideoNum: number
    MusicNum: number
    OtherNum: number
  }
}
export interface IndexResponse extends BasicResponse {
  data: OverallData
}
// docs, images, videos, audios, others页面响应
export interface DIVAOData {
  fileID: number
  fileName: string
  fileSize: number
  updateTime: string
  downloadNum: number
  fileUrl: string
}
export interface DIVAOResponse extends BasicResponse {
  data: DIVAOData[]
}

// 更换头像响应
export interface AvatarResponse extends BasicResponse {
  data: {
    avatar: string
  }
}
