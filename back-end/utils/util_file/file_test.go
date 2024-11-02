package util_file

import "testing"

func TestFileType(t *testing.T) {
	types := map[string]int{
		".doc": 1,
		".jpg": 2,
		".mp4": 3,
		".wav": 4,
		".ico": 5,
	}
	for k, v := range types {
		if GetFileType(k) != v {
			t.Fatal(k, "的类型判断错误，应该是", v, "，实际是", GetFileType(k))
		}
	}
}

func TestGetParentFolderNameAndPath(t *testing.T) {
	tests := map[string][2]string{
		"/go":      {"go", "/"},
		"/go/test": {"test", "/go"},
	}
	for k, v := range tests {
		if name, path := GetParentFolderNameAndPath(k); name != v[0] || path != v[1] {
			t.Fatal(k, "的父文件夹名称和路径判断错误，应该是", v, "，实际是", name, path)
		}
	}
}
