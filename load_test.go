package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	values, err := Load("fixtures/some.tfvars")

	assert.Nil(t, err)
	assert.Equal(t, len(values), 2, "we should have 2 return values")
}

func TestLoadConfigValues(t *testing.T) {
	values, _ := Load("fixtures/some.tfvars")

	assert.Equal(t, values[0].Name, "some_value")
	assert.Equal(t, values[0].Value, "a value")

	assert.Equal(t, values[1].Name, "some_ip")
	assert.Equal(t, values[1].Value, "172.20.50.0/25")
}
