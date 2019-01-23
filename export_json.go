package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func (ctx *Context) exportJSON() {
	data, err := jsoniter.MarshalToString(*ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}
