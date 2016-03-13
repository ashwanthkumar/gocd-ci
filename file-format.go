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
	return DecodeCISpecFromBytes([]byte(contents))
}

func DecodeCISpecFromBytes(data []byte) (CISpec, error) {
	ciSpec := CISpec{}
	err := yaml.Unmarshal(data, &ciSpec)
	if ciSpec.Cmd == nil {
		err = fmt.Errorf("Required field `cmd` is missing")
	}
	if ciSpec.Env == nil {
		ciSpec.Env = make(map[string]string)
	}
	return ciSpec, err
}
