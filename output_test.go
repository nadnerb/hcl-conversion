package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestOutput(t *testing.T) {

	all, err := Output([]*ConfigValue{
		&ConfigValue{
			Name: "a name",
			Value: "a value",
		},
	})

	assert.Nil(t, err)
	// this highlights an issue
	assert.Equal(t, all, "a name=a value")
}
