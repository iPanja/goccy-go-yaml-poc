package main

import (
	"os"

	"github.com/goccy/go-yaml"
)

func main() {
	b, err := os.ReadFile("input.yaml")
	if err != nil {
		panic(err)
	}

	print(string(b))

	// Parse the YAML file
	type datablock struct {
		A         *string            `yaml:"a"`
		Sequence  *[]string          `yaml:"seq"`
		Sequence2 *[]string          `yaml:"seq2"`
		Mapping   *map[string]string `yaml:"map"`
	}

	type doc struct {
		Src   datablock `yaml:"src"`
		Block datablock `yaml:"block"`
	}

	var data doc
	if err := yaml.Unmarshal(b, &data); err != nil {
		panic(err)
	}
	print("Parsed")

	// Test modifications
	//	These modifications unfortunately do not affect the `src` block, despite utilizing pointers
	new := "new"
	data.Block.A = &new
	(*data.Block.Sequence)[0] = "2"

	// Write the data back to a new YAML file
	b, err = yaml.Marshal(data)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("output.yaml", b, 0644); err != nil {
		panic(err)
	}
}
