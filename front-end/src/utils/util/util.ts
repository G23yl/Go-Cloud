// 工具函数

import type { IconInfo } from "@/types/types"

// 获取时间，返回morning，afternoon，evening
export function getTimeStr() {
  const time = new Date()
  const hour = time.getHours()
  if (hour < 12) {
    return "morning"
  } else if (hour < 18) {
    return "afternoon"
  } else {
    return "evening"
  }
}

// 转换文件大小
export function getFileSizeStr(fileSize: number) {
  if (fileSize < 1024) {
    return fileSize + " B"
  } else if (fileSize < 1024 * 1024) {
    return (fileSize / 1024).toFixed(2) + " KB"
  } else if (fileSize < 1024 * 1024 * 1024) {
    return (fileSize / 1024 / 1024).toFixed(2) + " MB"
  } else {
    return (fileSize / 1024 / 1024 / 1024).toFixed(2) + " GB"
  }
}

// 判断哪种文件类型确定对应icon
export function getFileTypeIcon(fileType: number): IconInfo {
  switch (fileType) {
    case 0:
      return { icon: ["fas", "folder"], color: "#ffd666" }
    case 1:
      return { icon: ["fas", "file-lines"], color: "#7f9bb8" }
    case 2:
      return { icon: ["fas", "image"], color: "#1388d3" }
    case 3:
      return { icon: ["fas", "film"], color: "#8d4be9" }
    case 4:
      return { icon: ["fas", "music"], color: "#dc7c73" }
    default:
      return { icon: ["fas", "circle-question"], color: "#dedfe0" }
  }
}
