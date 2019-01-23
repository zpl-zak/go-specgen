package gospec

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/inflection"
)

// ExportMD exports the data into a Markdown table format
func (ctx *Context) ExportMD() {
	for _, spec := range ctx.Specs {
		fmt.Printf("## %s\n\n", spec.Name)

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
