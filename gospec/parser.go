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

package gospec

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
	"strings"
)

// Field describes the data inside of spec
type Field struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	IsArray   bool   `json:"is_array"`
	IsPointer bool   `json:"is_ptr"`
	ArrayLen  uint   `json:"len"`
	DocString string `json:"_doc"`
}

// EnumField describes enum value
type EnumField struct {
	Value     string `json:"value"`
	DocString string `json:"_doc"`
}

// Enum describes enumeration
type Enum struct {
	Name   string      `json:"name"`
	Fields []EnumField `json:"fields"`
}

// Spec describes our structured data
type Spec struct {
	Name      string  `json:"name"`
	Fields    []Field `json:"fields"`
	DocString string  `json:"_doc"`
	resolved  bool
}

// Context contains all processed specs
type Context struct {
	FormatName string `json:"format"`
	Specs      []Spec `json:"specs"`
	Enums      []Enum `json:"enums"`
}

// ParseFile processes the provided gspec file
func ParseFile(filePath string) (Context, error) {
	ctx := Context{
		Specs: []Spec{},
		Enums: []Enum{},
	}

	fstFile := token.NewFileSet()
	node, err := parser.ParseFile(fstFile, filePath, nil, parser.ParseComments)

	if err != nil {
		fmt.Printf("Error parsing .gspec file: %v\n", err)
		return Context{}, fmt.Errorf("could not parse file: %v", err)
	}

	ctx.FormatName = node.Name.Name

	for _, decl := range node.Decls {
		spec, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		ctx.parseGenDecl(spec)
	}

	for _, comm := range node.Comments {
		txt := comm.Text()
		if idx := strings.Index(txt, "@enum"); idx != -1 {
			txt = txt[idx+6:]
			enumEnd := strings.Index(txt, "::")

			if enumEnd == -1 {
				continue
			}

			enumName := txt[:enumEnd]
			if enumName == "" {
				continue
			}

			txt = strings.TrimSpace(txt[enumEnd+2:])

			fields := []EnumField{}

			for {
				fld := strings.Index(txt, "->")

				if fld == -1 {
					break
				}

				txt = strings.TrimSpace(txt[fld+2:])
				comma := strings.Index(txt, ",")
				semicolon := strings.Index(txt, ";")

				if comma == -1 || comma > semicolon {
					val := txt[:semicolon]
					txt = strings.TrimSpace(txt[semicolon+1:])
					fields = append(fields, EnumField{
						Value: val,
					})
				} else if comma != -1 && comma < semicolon {
					val := txt[:comma]
					doc := txt[comma+1 : semicolon]
					txt = strings.TrimSpace(txt[semicolon+1:])
					fields = append(fields, EnumField{
						Value:     val,
						DocString: strings.TrimSpace(doc),
					})
				} else {
					break
				}
			}

			ctx.Enums = append(ctx.Enums, Enum{
				Name:   enumName,
				Fields: fields,
			})
		}
	}

	return ctx, nil
}

func (ctx *Context) parseGenDecl(decl *ast.GenDecl) {
	if len(decl.Specs) < 1 {
		return
	}

	spec, ok := decl.Specs[0].(*ast.TypeSpec)
	if !ok {
		return
	}

	name := spec.Name.Name
	doc := ""
	if spec.Comment != nil {
		doc = spec.Comment.Text()
	}

	specVal, ok := spec.Type.(*ast.StructType)
	if ok {
		ctx.parseSpec(name, doc, specVal)
		return
	}
}

func (ctx *Context) parseSpec(name, doc string, specVal *ast.StructType) {
	//spew.Dump(specVal)
	if specVal.Fields.NumFields() < 1 {
		fmt.Printf("Spec %s has no fields specified!\n", name)
		os.Exit(1)
		return
	}

	spec := Spec{
		Name:      name,
		Fields:    []Field{},
		DocString: strings.TrimSpace(doc),
	}

	for _, v := range specVal.Fields.List {
		var typeName string
		var isArray, isPtr bool
		var arrayLen uint
		var comment string
		typeVal, ok := v.Type.(*ast.Ident)
		if ok {
			typeName = typeVal.Name
		}
		arrayVal, ok := v.Type.(*ast.ArrayType)
		if ok {
			eltype, ok := arrayVal.Elt.(*ast.Ident)
			if !ok {
				fmt.Printf("Field %s in spec %s at %d can't be array of arrays!\n", v.Names[0].Name, name, arrayVal.Pos())
				os.Exit(2)
				return
			}
			typeName = eltype.Name

			lenVal, ok := arrayVal.Len.(*ast.BasicLit)
			if !ok {
				arrayLen = 0
			} else {
				lenConv, _ := strconv.Atoi(lenVal.Value)
				arrayLen = uint(lenConv)
			}
			isArray = true
		}

		if v.Comment != nil {
			comment = v.Comment.Text()
		}

		if strings.Contains(comment, "@ptr") {
			isPtr = true
		}

		for _, name := range v.Names {
			field := Field{
				Name:      name.Name,
				Type:      typeName,
				IsArray:   isArray,
				IsPointer: isPtr,
				ArrayLen:  arrayLen,
				DocString: comment,
			}

			spec.Fields = append(spec.Fields, field)
		}
	}

	ctx.Specs = append(ctx.Specs, spec)
}
