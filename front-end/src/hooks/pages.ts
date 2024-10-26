// 获取数据

import { indexReq } from "@/api/page"

export const usePageData = () => {
  // index页面数据获取
  async function index() {
    const response = await indexReq()
    return response
  }

  return { index }
}
