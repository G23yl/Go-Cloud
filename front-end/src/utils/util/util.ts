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
