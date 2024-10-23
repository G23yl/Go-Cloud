// 定义数据结构

export interface LoginForm {
  username: string
  password: string
}

// 基础响应包含code和msg，但是data字段根据具体情况返回对应类型
interface BasicResponse {
  code: number
  msg: string
}

export interface User {
  avatar: string
  email: string
  token: string
  username: string
}

// 继承BasicResponse，根据具体返回情况设置data字段的类型
export interface LoginResponse extends BasicResponse {
  data: User
}

// 菜单条目
export interface MenuItem {
  index: string
  itemName: string
}
export type MenuItems = MenuItem[]
