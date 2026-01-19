package jsonutil

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func GetField(data []byte, key string) (string, error) {
	var result string

	err := Process(data, func(v *fastjson.Value) error {
		val := v.Get(key)
		if val == nil {
			return fmt.Errorf("key '%s' not found", key)
		}

		result = string(val.GetStringBytes())
		return nil
	})

	return result, err
}
