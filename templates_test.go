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

func Test_GetFilesForEmptyInput(t *testing.T) {
	expected := []TemplateFile{}

	assert.Equal(t, expected, ParseInput(""))
}

func Test_GetFilesForValidSingleInput(t *testing.T) {
	expected := []TemplateFile{"Go"}

	assert.Equal(t, expected, ParseInput("go"))
}

func Test_GetFilesForSameInputDuplicated(t *testing.T) {
	expected := []TemplateFile{"Go"}

	assert.Equal(t, expected, ParseInput("go golang"))
}

func Test_GetFilesForMultipleInput(t *testing.T) {
	expected := []TemplateFile{"Go", "JavaScript"}

	assert.ElementsMatch(t, expected, ParseInput("go js"))
}
