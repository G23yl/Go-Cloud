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
