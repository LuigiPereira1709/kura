package fsutil

import (
	"fmt"
	"os"
	"path/filepath"
)

func DeleteFilesFromDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", dir, err)
	}

	for _, file := range files {
		if err := os.RemoveAll(filepath.Join(dir, file.Name())); err != nil {
			return fmt.Errorf("failed to delete file %s: %w", file.Name(), err)
		}
	}
	return nil
}
