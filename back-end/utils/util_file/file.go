// 定义关于文件相关操作的工具函数
package util_file

import (
	"os"
	"strings"
)

// 获取某个文件夹下所有文件
func GetAllFiles(path string) ([]string, error) {
	var files []string
	entry, err := os.ReadDir(path)
	if err != nil {
		return files, err
	}
	for _, e := range entry {
		if !e.IsDir() {
			files = append(files, e.Name())
		}
	}
	return files, nil
}

// 判断文件类型
func GetFileType(suffix string) int {
	fileSuffix := strings.ToLower(suffix)
	switch fileSuffix {
	case ".doc":
		fallthrough
	case ".docx":
		fallthrough
	case ".txt":
		fallthrough
	case ".pdf":
		return 1

	case ".jpg":
		fallthrough
	case ".png":
		fallthrough
	case ".gif":
		fallthrough
	case ".jpeg":
		return 2

	case ".mp4":
		fallthrough
	case ".avi":
		fallthrough
	case ".mov":
		fallthrough
	case ".rmvb":
		fallthrough
	case ".rm":
		return 3

	case ".mp3", ".cda", ".wav", ".wma", ".ogg":
		return 4
	default:
		return 5
	}
}
