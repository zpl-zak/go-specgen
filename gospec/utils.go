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
)

func retrieveSortedSpecs(specs []Spec) (out []Spec) {
	out = []Spec{}
	for _, spec := range specs {
		if spec.resolved {
			continue
		}

		resolveDeps(specs, &out, spec, spec)
		out = append(out, spec)
	}

	//spew.Dump(out)
	return out
}

func resolveDeps(specs []Spec, out *[]Spec, spec, orig Spec) {
	for _, field := range spec.Fields {
		for i, dep := range specs {
			if field.Type == dep.Name {
				if dep.Name == orig.Name {
					fmt.Printf("Error: Cyclic dependency used for field %s in spec %s!\n", field.Name, spec.Name)
					return
				}

				specs[i].resolved = true
				resolveDeps(specs, out, dep, orig)
				*out = append(*out, dep)
			}
		}
	}
}