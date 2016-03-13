package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type CISpec struct {
	Name      string            `yaml:"name"`
	Env       map[string]string `yaml:"env,omitempty"`
	Cmd       []string          `yaml:"cmd"`
	Artifacts map[string]string `yaml:"artifacts,omitempty"`
}

func DecodeCISpec(contents string) (CISpec, error) {
	ciSpec := CISpec{}
	err := yaml.Unmarshal([]byte(contents), &ciSpec)
	if ciSpec.Cmd == nil {
		err = fmt.Errorf("Required field `cmd` is missing")
	}
	return ciSpec, err
}
