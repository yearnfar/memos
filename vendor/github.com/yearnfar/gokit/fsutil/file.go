package fsutil

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// IsFile 检查是否文件
func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	if err == nil && !info.IsDir() {
		return true
	}
	return false
}

// FileName 只返回文件名称
func FileName(fp string) string {
	name := filepath.Base(fp)
	suffix := filepath.Ext(name)
	return strings.TrimSuffix(name, suffix)
}

// FileExt 返回文件扩展名，如 foo.png 返回 .png
func FileExt(path string) string {
	return strings.ToLower(filepath.Ext(path))
}

// FileSize 返回文件大小
func FileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

// HumanReadableSize 可读大小
func HumanReadableSize(n int64) string {
	const unit = 1024
	i, size := 0, float64(n)
	for size >= unit {
		size /= unit
		i++
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return fmt.Sprintf("%.2f %s", size, units[i])
}

// MD5File md5文件
func MD5File(fn string) (string, error) {
	file, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	defer file.Close()
	const bufferSize = 10 << 20 // 10M
	buffer := make([]byte, bufferSize)
	reader := bufio.NewReader(file)
	hash := md5.New()
	for {
		n, err := reader.Read(buffer)
		if err == nil || err == io.EOF {
			_, _ = hash.Write(buffer[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
