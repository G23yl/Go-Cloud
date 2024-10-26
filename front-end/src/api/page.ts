// 页面数据获取api

import type { IndexResponse } from "@/types/types"
import request from "@/utils/axios"

export const indexReq = async () => {
  try {
    const res = await request.get<any, IndexResponse>("/cloud/index")
    return res
  } catch (error) {
    console.log(error)
  }
}
