package jsonutil

import (
	"fmt"

	json "github.com/goccy/go-json"
)

func ToMap(data []byte) (map[string]any, error) {
	var result map[string]any
	if err := json.UnmarshalNoEscape(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return result, nil
}

