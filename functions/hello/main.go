package main

import (
	"github.com/apex/go-apex"
	"encoding/json"
)


func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		//やりたい処理
		hoge := "Hello World"
		return hoge, nil
	})
}
