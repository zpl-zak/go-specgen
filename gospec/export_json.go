package gospec

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

// ExportJSON marshals our current data into JSON format
func (ctx *Context) ExportJSON() {
	data, err := jsoniter.MarshalToString(*ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}
