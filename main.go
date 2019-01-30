/*
   Copyright 2019 Dominik Madarász
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at
       http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/zaklaus/go-specgen/gospec"
)

func main() {
	filePath := flag.String("file", "", "gspec file to generate from")
	langMode := flag.String("lang", "json", "language mode to use")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Not enough arguments!\n")
		flag.PrintDefaults()
		os.Exit(3)
		return
	}

	if _, err := os.Stat(*filePath); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "File does not exist!")
		os.Exit(4)
		return
	}

	ctx, err := gospec.ParseFile(*filePath)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	switch *langMode {
	case "json":
		ctx.ExportJSON()
	case "c":
		ctx.ExportC()
	case "md":
		ctx.ExportMD()
	case "go":
		ctx.ExportGo()
	case "dump":
		spew.Dump(ctx)

	default:
		fmt.Fprintf(os.Stderr, "Lang mode not supported!\n")
		os.Exit(5)
		return
	}
}
