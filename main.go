package main

import (
	"fmt"
)

func main() {
	ctx, err := parseFile("drafts/foo.gspec")

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	ctx.exportC()
	//spew.Dump(ctx)
}
