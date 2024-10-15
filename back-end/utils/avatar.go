package utils

import "disk/config"

// 获取头像url地址
func GetImageUrl(filename string) string {
	return "http://" + config.GetServerHost() + config.GetServerPort() + "/images/" + filename
}
