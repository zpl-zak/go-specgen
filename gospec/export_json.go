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
	"os"

	jsoniter "github.com/json-iterator/go"
)

// ExportJSON marshals our current data into JSON format
func (ctx *Context) ExportJSON() {
	data, err := jsoniter.MarshalToString(*ctx)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}

	fmt.Println(data)
}
