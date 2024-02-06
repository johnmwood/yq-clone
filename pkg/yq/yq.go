package yq

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Record struct {
	yaml.Node
}

func ParseYAML(filepath string) (Record, error) {
	if filepath[len(filepath)-5:] != ".yaml" {
		return Record{}, fmt.Errorf("file is not a yaml file")
	}
	data, err := os.ReadFile(filepath)
	if err != nil {
		return Record{}, fmt.Errorf("error reading file: %v", err)
	}

	var record Record
	err = yaml.Unmarshal(data, &record.Node)
	if err != nil {
		return Record{}, fmt.Errorf("error building yaml graph: %v", err)
	}
	return record, nil
}
