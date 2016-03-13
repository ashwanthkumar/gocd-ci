package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeCISpec(t *testing.T) {
	var sampleSpec = `
name: "GoCD-CI"
env:
  FOO: bar
  BAR: baz
  BAZ: day
cmd:
  - echo "Hello World"
artifacts: {}
`
	spec, err := DecodeCISpec(sampleSpec)
	assert.Nil(t, err)
	expectedSpec := CISpec{
		Name:      "GoCD-CI",
		Env:       map[string]string{"FOO": "bar", "BAR": "baz", "BAZ": "day"},
		Cmd:       []string{"echo \"Hello World\""},
		Artifacts: make(map[string]string),
	}

	assert.Equal(t, expectedSpec, spec)
}

func TestDecodeCISpecWithoutArtifacts(t *testing.T) {
	var sampleSpec = `
name: "GoCD-CI"
env:
  FOO: bar
  BAR: baz
  BAZ: day
cmd:
  - echo "Hello World"
`
	spec, err := DecodeCISpec(sampleSpec)
	assert.Nil(t, err)
	expectedSpec := CISpec{
		Name: "GoCD-CI",
		Env:  map[string]string{"FOO": "bar", "BAR": "baz", "BAZ": "day"},
		Cmd:  []string{"echo \"Hello World\""},
	}

	assert.Equal(t, expectedSpec, spec)
}

func TestDecodeCISpecWithoutEnv(t *testing.T) {
	var sampleSpec = `
name: "GoCD-CI"
cmd:
  - echo "Hello World"
`
	spec, err := DecodeCISpec(sampleSpec)
	assert.Nil(t, err)
	expectedSpec := CISpec{
		Name: "GoCD-CI",
		Cmd:  []string{"echo \"Hello World\""},
	}

	assert.Equal(t, expectedSpec, spec)
}

func TestDecodeCISpecWithoutCommandShouldError(t *testing.T) {
	var sampleSpec = `
name: "GoCD-CI"
env:
  FOO: bar
  BAR: baz
  BAZ: day
`
	_, err := DecodeCISpec(sampleSpec)
	assert.Error(t, err)
}
