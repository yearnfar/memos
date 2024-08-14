package fsutil

import "os"

// IsDir 检查是否目录
func IsDir(dir string) bool {
	info, err := os.Stat(dir)
	if err == nil && info.IsDir() {
		return true
	}
	return false
}
