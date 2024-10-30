// 工具函数

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
