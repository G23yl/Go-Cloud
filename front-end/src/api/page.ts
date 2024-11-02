// 页面数据获取api

import type { BasicResponse, DIVAOResponse, FFResponse, IndexResponse } from "@/types/types"
import request from "@/utils/axios"

// index页面数据api
export const indexReq = async () => {
  try {
    const res = await request.get<any, IndexResponse>("/cloud/index")
    return res
  } catch (error) {
    console.log(error)
  }
}

// docs页面数据api
export const docsReq = async () => {
  try {
    const res = await request.get<any, DIVAOResponse>("/cloud/docs")
    return res
  } catch (error) {
    console.log(error)
  }
}

// images页面数据api
export const imagesReq = async () => {
  try {
    const res = await request.get<any, DIVAOResponse>("/cloud/images")
    return res
  } catch (error) {
    console.log(error)
  }
}

// videos页面数据api
export const videosReq = async () => {
  try {
    const res = await request.get<any, DIVAOResponse>("/cloud/videos")
    return res
  } catch (error) {
    console.log(error)
  }
}

// audios页面数据api
export const audiosReq = async () => {
  try {
    const res = await request.get<any, DIVAOResponse>("/cloud/audios")
    return res
  } catch (error) {
    console.log(error)
  }
}

// others页面数据api
export const othersReq = async () => {
  try {
    const res = await request.get<any, DIVAOResponse>("/cloud/others")
    return res
  } catch (error) {
    console.log(error)
  }
}

// 总页面数据api
export const filesReq = async (path: string) => {
  try {
    const res = await request.get<any, FFResponse>(`/cloud/files`, {
      params: {
        path: path,
      },
    })
    return res
  } catch (error) {
    console.log(error)
  }
}

// 删除文件
export const deleteFReq = async (filePath: string, fileName: string, fileID: number) => {
  try {
    const res = await request.delete<any, BasicResponse>("/cloud/files", {
      params: {
        fileID: fileID,
        filePath: filePath,
        fileName: fileName,
      },
    })
    return res
  } catch (error) {
    console.log(error)
  }
}

// 删除文件夹
export const deleteFolderReq = async (folderPath: string, folderName: string, folderID: number) => {
  try {
    const res = await request.delete<any, BasicResponse>("cloud/folders", {
      params: {
        folderID: folderID,
        folderPath: folderPath,
        folderName: folderName,
      },
    })
    return res
  } catch (error) {
    console.log(error)
  }
}

// 创建文件夹
export const createFolderReq = async (path: string, folderName: string) => {
  try {
    const res = await request.post<any, BasicResponse>(
      "/cloud/folders",
      { folderName: folderName },
      {
        params: {
          path: path,
        },
      }
    )
    return res
  } catch (error) {
    console.log(error)
  }
}

// 下载文件
export const downloadFileReq = async (filePath: string, fileName: string) => {
  try {
    const res = await request.get<any, any>("/cloud/download", {
      params: {
        filePath: filePath,
        fileName: fileName,
      },
      responseType: "blob",
    })
    return res
  } catch (error) {
    console.log(error)
  }
}
