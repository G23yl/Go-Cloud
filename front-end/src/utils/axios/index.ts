// 二次封装axios

import type { LoginResponse } from "@/types/types"
import axios, { type AxiosError } from "axios"
import { ElMessage } from "element-plus"
import { getUser } from "../user"

const request = axios.create({
  timeout: 5000,
  // baseURL: import.meta.env.VITE_APP_BASE_URL as string,
  baseURL: "/api",
})

request.interceptors.request.use(
  (config) => {
    // 为后续请求添加token，如果有token表示用户已登录，如果没token表示用户退出登录，处于未登录状态
    config.headers.Authorization = getUser()?.token
    return config
  },
  (error: AxiosError) => {
    return Promise.reject(error)
  }
)

request.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error: AxiosError) => {
    // 网络超时
    if (error.message === "Network Error" || error.code === "ERR_NETWORK") {
      ElMessage({
        type: "error",
        message: "网络错误",
      })
    }
    // 非2xx响应提示
    if (error.response && error.response.data) {
      ElMessage({
        type: "error",
        message: (error.response.data as LoginResponse).msg,
      })
    }
    return Promise.reject(error)
  }
)

export default request
