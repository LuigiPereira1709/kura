package fsutil

import (
	"fmt"
	"os"
	"strings"
)

func CreateDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func GetParentDir(dir string) string {
	// TODO: Improve way to get the parent dir
	idx := strings.LastIndex(dir, "/")
	if idx == -1 {
		return ""
	}

	parentDir := dir[:idx]
	return parentDir
}

func OpenFile(filePath string) (*os.File, error) {
	if filePath == "" {
		return nil, fmt.Errorf("file path cannot be empty")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}

	return file, nil
}
