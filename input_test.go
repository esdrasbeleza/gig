package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Input_Lowercase(t *testing.T) {
	var (
		input  = Input("Go")
		output = Input("go")
	)

	assert.Equal(t, output, input.Lowercase())
}
