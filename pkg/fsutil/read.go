package fsutil

import (
	"fmt"
	"io"
)

func Read(filePath string) ([]byte, error) {
	file, err := OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get stats from file: %v", err)
	}

	buf := make([]byte, stats.Size())
	_, err = io.ReadFull(file, buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	return buf, nil
}
