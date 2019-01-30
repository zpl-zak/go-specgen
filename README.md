# go-specgen

Go specgen is a minimalistic library for generating format specification into various languages.

It uses Go semantics (see `drafts/foo.gspec`) and currently generates the specifications to various formats.

## Things to do
- Add more languages/formats
    - [x] C/C++
    - [X] Go
    - [X] JSON
    - [X] Markdown
    - [ ] Python
    - [ ] Rust
    - [ ] C#
    - [ ] Java
- [x] Improve error handling
- [x] Implement enum types

## How to use

Simply `go get github.com/zaklaus/go-specgen/gospec` or clone the repository.

### Usage: Tool
You can compile go-specgen and use it as a tool to generate various outputs easily.

You can generate output to supported language by doing:
```sh
specgen --file=<path-to-file> --lang=<lang-mode>
```

such as 
```sh
specgen --file=drafts/foo.gspec --lang=c
```

which will print out the output to the stdout stream. 

For instance, you could pipe the output of Go lang into `gofmt` to get nicely formatted Go source code:
```sh
specgen --file=drafts/foo.gspec --lang=go | gofmt
```

### Usage: Library
You can also use go-specgen as a library, for instance:
```go
package main

import (
    "fmt"

    "github.com/zaklaus/go-specgen/gospec"
)

func main() {
    // Parse the gspec file containing data specifications
    ctx, err := gospec.ParseFile("drafts/foo.gspec")

    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Export it to Markdown tables
    ctx.ExportMD()
}
```

You can also look at `main.go` itself.

## License

This project is licensed under Apache 2.0, see LICENSE.md
