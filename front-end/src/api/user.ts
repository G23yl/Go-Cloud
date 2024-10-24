// 发起请求的接口

import type { BasicResponse, LoginForm, LoginResponse, SignUpForm } from "@/types/types"
import request from "@/utils/axios"

export const loginReq = async (data: LoginForm) => {
  try {
    const res = await request.post<any, LoginResponse>("/login", data)
    return res
  } catch (error) {
    console.log(error)
  }
}

export const sendCodeReq = async (email: string) => {
  try {
    const res = await request.post<any, BasicResponse>("/verify", { email: email })
    return res
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
    const res = await request.post<any, BasicResponse>("/register", formData)
    return res
  } catch (error) {
    console.log(error)
  }
}
