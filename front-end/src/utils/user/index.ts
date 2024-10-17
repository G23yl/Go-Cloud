// 用户信息操作

import type { User } from "@/types/types"

function getUser() {
  return JSON.parse(localStorage.getItem("user") || "")
}

function setUser(user: User) {
  localStorage.setItem("user", JSON.stringify(user))
}

function deleteUser() {
  localStorage.removeItem("user")
}

export { getUser, setUser, deleteUser }
