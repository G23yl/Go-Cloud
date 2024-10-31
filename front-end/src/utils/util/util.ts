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

// 判断哪种文件类型确定对应icon
export function getFileTypeIcon(fileType: number) {
  switch (fileType) {
    case 0:
      return ["fas", "folder"]
    case 1:
      return ["fas", "file"]
    case 2:
      return ["fas", "image"]
    case 3:
      return ["fas", "film"]
    case 4:
      return ["fas", "music"]
    default:
      return ["fas", "circle-question"]
  }
}
