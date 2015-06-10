package main

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
	hclobj "github.com/hashicorp/hcl/hcl"
)

type ConfigValue struct {
	Name						string
	Value						string
}

func Load(path string) ([]*ConfigValue, error) {
	// Read the HCL file and prepare for parsing
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(
			"Error reading %s: %s", path, err)
	}

	// Parse it
	obj, err := hcl.Parse(string(file))

	var all []*hclobj.Object
	// probably a better way to do this, need to test multiple formats of hcl files
	// i ripped this out of terraform, need to try and understand it...
	for _, o1 := range obj.Elem(false) {
		// Iterate the inner to get the list of types
		for _, o2 := range o1.Elem(true) {
			// Iterate all of this type to get _all_ the types
			for _, o3 := range o2.Elem(false) {
				all = append(all, o3)
			}
		}
	}

	var result []*ConfigValue
	for _, obj := range all {
		result = append(result, &ConfigValue{
			Name: obj.Key,
			Value: obj.Value.(string),
		})
	}

	if err != nil {
		return nil, fmt.Errorf(
			"Error parsing %s: %s", path, err)
	}

	// Build up the result
	//var result Config
	//if err := hcl.DecodeObject(&result, obj); err != nil {
		//return nil, err
	//}

	return result, nil
}

