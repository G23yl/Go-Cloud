// 用户信息操作

import type { User } from "@/types/types"

function getUser() {
  let user = localStorage.getItem("user")
  if (user) {
    return JSON.parse(user)
  }
  return {}
}

function setUser(user: User) {
  localStorage.setItem("user", JSON.stringify(user))
}

function deleteUser() {
  localStorage.removeItem("user")
}

export { getUser, setUser, deleteUser }
