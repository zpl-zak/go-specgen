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
	"strconv"
	"strings"

	"github.com/jinzhu/inflection"
)

// ExportMD exports the data into a Markdown table format
func (ctx *Context) ExportMD() {
	for _, enum := range ctx.Enums {
		fmt.Printf("## Enum: %s\n\n", enum.Name)

		fmt.Println("| Value | Description |")
		fmt.Println("| ----- | ----------- |")

		for _, field := range enum.Fields {
			fmt.Printf("| %s | %s |\n", field.Value, field.DocString)
		}
	}

	fmt.Println("")

	for _, spec := range ctx.Specs {
		fmt.Printf("## Spec: %s\n\n", spec.Name)

		fmt.Println("| Type | Name | Description |")
		fmt.Println("| ---- | ---- | ----------- |")

		for _, field := range spec.Fields {
			fmt.Printf("| %s | %s | %s%s |\n", dumpType(field), field.Name, dumpSpecials(&field), strings.TrimSpace(field.DocString))
		}
	}

	fmt.Println("")
}

func dumpSpecials(field *Field) string {
	if field.IsArray {
		lenName := strconv.Itoa(int(field.ArrayLen))
		if lenName == "0" {
			lenName = "N"
		}

		singName := inflection.Singular(field.Name)

		if strings.Contains(field.DocString, "@string") {
			field.DocString = strings.Replace(field.DocString, "@string", "", -1)
			return fmt.Sprintf("String consisting of %s characters; ", lenName)
		}

		if strings.Contains(field.DocString, "@plain") {
			field.DocString = strings.Replace(field.DocString, "@plain", "", -1)
			return fmt.Sprintf("plain array of %s elements; ", lenName)
		}

		return fmt.Sprintf("%s definitions of %s; ", lenName, singName)
	}

	return ""
}

func dumpType(field Field) string {
	if field.IsPointer {
		return fmt.Sprintf("%s*", field.Type)
	}

	return field.Type
}
