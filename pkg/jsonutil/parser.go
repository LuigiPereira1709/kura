package jsonutil

import (
	"encoding/json"
	"fmt"
)

func ToMap(data []byte) (map[string]any, error) {
	// TODO: Improve this Function to low level approach
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return result, nil
}
