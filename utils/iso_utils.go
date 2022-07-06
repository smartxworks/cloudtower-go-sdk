package utils

import (
	"os"
	"strings"
)

func CheckSvtImageVersion(file *os.File) (string, error) {
	bytes := make([]byte, 128)
	_, err := file.ReadAt(bytes, 32*1024+190)
	if err != nil {
		return "", err
	}
	version := string(bytes)
	return strings.TrimSpace(version), nil
}
