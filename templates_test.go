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

func Test_GetTemplateThatDoesNotExist(t *testing.T) {
	assert.Equal(t, TemplateFile(""), GetTemplate("Weird programming language"))
}
