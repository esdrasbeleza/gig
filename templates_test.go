package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTemplate(t *testing.T) {
	var (
		input    = []Input{"go", "Go", "golang"}
		expected = TemplateFile("Go")
	)

	for _, i := range input {
		assert.Equal(t, expected, GetTemplate(i))
	}
}
