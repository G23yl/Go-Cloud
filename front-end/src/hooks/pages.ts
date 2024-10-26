// 获取数据

import { audiosReq, docsReq, imagesReq, indexReq, othersReq, videosReq } from "@/api/page"

export const usePageData = () => {
  // index页面数据获取
  async function index() {
    return await indexReq()
  }
  // docs页面数据获取
  async function docs() {
    return await docsReq()
  }
  // images页面数据获取
  async function images() {
    return await imagesReq()
  }
  // videos页面数据获取
  async function videos() {
    return await videosReq()
  }
  // audios页面数据获取
  async function audios() {
    return await audiosReq()
  }
  // others页面数据获取
  async function others() {
    return await othersReq()
  }

  return { index, docs, images, videos, audios, others }
}
