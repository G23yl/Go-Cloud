// 用户仓库

import { loginReq } from "@/api/user"
import type { LoginForm, User } from "@/types/types"
import { getUser, setUser } from "@/utils/user"
import { acceptHMRUpdate, defineStore } from "pinia"

const useUserStore = defineStore("user", () => {
  let user: User = getUser()

  async function login(data: LoginForm) {
    const response = await loginReq(data)
    // 登录成功后，将user存储到localStorage中
    // console.log(response)
    if (response) {
      user = response.data
      setUser(user)
    }
    return response
  }

  return { login, user }
})

// 热重载
if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useUserStore, import.meta.hot))
}
export default useUserStore
