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

func Test_FileForLanguage(t *testing.T) {
	var (
		input    = []Input{"go", "Go", "golang"}
		expected = TemplateFile("Go")
	)

	for _, i := range input {
		assert.Equal(t, expected, FileForLanguage(i))
	}
}
