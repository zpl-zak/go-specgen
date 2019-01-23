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

// Spec describes our structured data
type Spec struct {
	Name      string  `json:"name"`
	Fields    []Field `json:"fields"`
	DocString string  `json:"_doc"`
	resolved  bool
}

// Context contains all processed specs
type Context struct {
	Specs []Spec `json:"specs"`
	// TODO(zaklaus): Methods
}

// ParseFile processes the provided gspec file
func ParseFile(filePath string) (Context, error) {
	ctx := Context{
		Specs: []Spec{},
	}

	fstFile := token.NewFileSet()
	node, err := parser.ParseFile(fstFile, filePath, nil, parser.ParseComments)

	if err != nil {
		fmt.Printf("Error parsing .gspec file: %v\n", err)
		return Context{}, fmt.Errorf("could not parse file: %v", err)
	}

	for _, decl := range node.Decls {
		spec, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		ctx.parseGenDecl(spec)
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

	specVal, ok := spec.Type.(*ast.StructType)
	if ok {
		ctx.parseSpec(name, specVal)
		return
	}
}

func (ctx *Context) parseSpec(name string, specVal *ast.StructType) {
	//spew.Dump(specVal)
	if specVal.Fields.NumFields() < 1 {
		fmt.Printf("Spec %s has no fields specified!\n", name)
		os.Exit(1)
		return
	}

	spec := Spec{
		Name:   name,
		Fields: []Field{},
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

		if strings.Contains(typeName, "ptr") {
			typeName = strings.Replace(typeName, "ptr", "", -1)
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
