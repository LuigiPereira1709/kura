package fsutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func WriteFileFromReader(dir, fileName string, reader io.Reader) (string, error) {
	if fileName == "" {
		return "", fmt.Errorf("file name cannot be empty")
	}

	if err := CreateDir(dir); err != nil {
		return "", fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	filePath := filepath.Join(dir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	if _, err := io.Copy(file, reader); err != nil { return "", fmt.Errorf("failed to write to file %s: %w", filePath, err)
	}
	return filePath, file.Sync()
}
