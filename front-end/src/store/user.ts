// 用户仓库

import { loginReq, sendCodeReq, signUpReq } from "@/api/user"
import type { LoginForm, SignUpForm } from "@/types/types"
import { setUser } from "@/utils/user"
import { acceptHMRUpdate, defineStore } from "pinia"

const useUserStore = defineStore("user", () => {
  async function login(data: LoginForm) {
    const response = await loginReq(data)
    // 登录成功后，将user存储到localStorage中
    // console.log(response)
    if (response) {
      setUser(response.data)
    }
    return response
  }

  // 发送验证码
  async function sendCode(email: string) {
    return await sendCodeReq(email)
  }

  // 注册
  async function signUp(data: SignUpForm) {
    return await signUpReq(data)
  }

  return { login, sendCode, signUp }
})

// 热重载
if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useUserStore, import.meta.hot))
}
export default useUserStore
