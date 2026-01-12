package jsonutil

import (
	"fmt"

	"github.com/valyala/fastjson"
)

var pool fastjson.ParserPool

func Process(data []byte, handle func(*fastjson.Value) error) error {
	p := pool.Get()
	defer pool.Put(p)

	v, err := p.ParseBytes(data)
	if err != nil {
		return fmt.Errorf("failed to parse data: %s", err)
	}

	return handle(v)
}
