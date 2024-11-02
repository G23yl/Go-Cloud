// 定义关于文件相关操作的工具函数
package util_file

import (
	"os"
	"slices"
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

// 各种类型的数组集合
var (
	DocType   = []string{".doc", ".docx", ".txt", ".pdf", ".md", ".xls"}
	ImgType   = []string{".jpg", ".png", ".gif", ".jpeg", ".svg"}
	VideoType = []string{".mp4", ".avi", ".mov", ".rmvb", ".rm"}
	AudioType = []string{".mp3", ".cda", ".wav", ".wma", ".ogg", ".m4a"}
)

// 判断文件类型
func GetFileType(suffix string) int {
	fileSuffix := strings.ToLower(suffix)
	if slices.Contains(DocType, fileSuffix) {
		return 1
	} else if slices.Contains(ImgType, fileSuffix) {
		return 2
	} else if slices.Contains(VideoType, fileSuffix) {
		return 3
	} else if slices.Contains(AudioType, fileSuffix) {
		return 4
	} else {
		return 5
	}
}

func GetParentFolderNameAndPath(filePath string) (string, string) {
	paths := strings.Split(filePath, "/")
	parentFolderName := paths[len(paths)-1]
	paths = paths[:len(paths)-1]
	if len(paths) == 1 {
		return parentFolderName, "/"
	}
	parentFolderPath := strings.Join(paths, "/")
	return parentFolderName, parentFolderPath
}
