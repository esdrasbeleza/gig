package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTemplate(t *testing.T) {
	var (
		input    = []Key{"go", "Go", "golang"}
		expected = Key("Go")
	)

	for _, i := range input {
		template := GetTemplate(i)
		assert.NotNil(t, template, "Got nil template for %s", i)
		assert.Equal(t, expected, template.Name)
	}
}

func Test_GetTemplateThatDoesNotExist(t *testing.T) {
	assert.Nil(t, GetTemplate("Weird programming language"))
}

func Test_GetFilesForEmptyInput(t *testing.T) {
	expected := []Key{}

	assert.Equal(t, expected, ParseInput("").Keys())
}

func Test_GetFilesForValidSingleInput(t *testing.T) {
	expected := []Key{"Go"}

	assert.Equal(t, expected, ParseInput("go").Keys())
}

func Test_GetFilesForValidSingleInputVariation(t *testing.T) {
	expected := []Key{"Go"}

	assert.Equal(t, expected, ParseInput("Go").Keys())
}

func Test_GetFilesForSameInputDuplicated(t *testing.T) {
	expected := []Key{"Go"}

	assert.Equal(t, expected, ParseInput("go golang").Keys())
}

func Test_GetNoFilesForInvalidInput(t *testing.T) {
	expected := []Key{}

	assert.Equal(t, expected, ParseInput("asdf").Keys())
}

func Test_GetFilesForMultipleInput(t *testing.T) {
	expected := []Key{"Go", "VisualStudioCode"}

	assert.ElementsMatch(t, expected, ParseInput("go code").Keys())
}
