// 发起请求的接口

import type { LoginForm, LoginResponse } from "@/types/types"
import request from "@/utils/axios"

export const loginReq = async (data: LoginForm) => {
  try {
    const res = await request.post<any, LoginResponse>("/login", data)
    return res
  } catch (error) {
    console.log(error)
  }
}
