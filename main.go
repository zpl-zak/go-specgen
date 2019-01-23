package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/zaklaus/go-specgen/gospec"
)

func main() {
	filePath := flag.String("file", "test.gspec", "gspec file to generate from")
	langMode := flag.String("lang", "json", "language mode to use")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("Not enough arguments!\n")
		flag.PrintDefaults()
		os.Exit(3)
		return
	}

	if _, err := os.Stat(*filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist!")
		os.Exit(4)
		return
	}

	ctx, err := gospec.ParseFile(*filePath)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	switch *langMode {
	case "json":
		ctx.ExportJSON()
	case "c":
		ctx.ExportC()
	case "md":
		ctx.ExportMD()
	case "dump":
		spew.Dump(ctx)

	default:
		fmt.Printf("Lang mode not supported!\n")
		os.Exit(5)
		return
	}
}
