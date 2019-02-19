package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Input_Lowercase(t *testing.T) {
	var (
		input  = Key("Go")
		output = Key("go")
	)

	assert.Equal(t, output, input.Lowercase())
}
