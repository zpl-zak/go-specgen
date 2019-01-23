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
