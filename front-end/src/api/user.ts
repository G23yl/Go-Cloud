// 发起请求的接口

import type {
  AvatarResponse,
  BasicResponse,
  LoginForm,
  LoginResponse,
  SignUpForm,
} from "@/types/types"
import request from "@/utils/axios"

export const loginReq = async (data: LoginForm) => {
  try {
    return await request.post<any, LoginResponse>("/login", data)
  } catch (error) {
    console.log(error)
  }
}

export const sendCodeReq = async (email: string) => {
  try {
    return await request.post<any, BasicResponse>("/verify", { email: email })
  } catch (error) {
    console.log(error)
  }
}

export const signUpReq = async (data: SignUpForm) => {
  try {
    const formData = new FormData()
    formData.append("username", data.username)
    formData.append("password", data.password)
    formData.append("code", data.code)
    formData.append("email", data.email)
    return await request.post<any, BasicResponse>("/register", formData)
  } catch (error) {
    console.log(error)
  }
}

export const changeAvatarReq = async (data: File) => {
  try {
    const formData = new FormData()
    formData.append("avatar", data)
    return await request.post<any, AvatarResponse>("/user/avatar", formData)
  } catch (error) {
    console.log(error)
  }
}
