package main

import (
	"fmt"
	"strings"
)

var (
	inputAliasesMap = map[TemplateName][]Input{
		"Go":         {"Go", "golang"},
		"JavaScript": {"JavaScript", "JS"},
		"TypeScript": {"TypeScript", "TS"},
	}

	templateMap map[Input]TemplateName
)

func generateLanguageMaps() {
	templateMap = make(map[Input]TemplateName)

	for templateName, templateAliases := range inputAliasesMap {
		for _, templateAlias := range templateAliases {
			templateMap[templateAlias.Lowercase()] = templateName
		}
	}
}

func GetTemplate(input Input) TemplateName {
	if languageFile, exists := templateMap[input.Lowercase()]; exists {
		return languageFile
	}

	return TemplateName("")
}

func ParseInput(input string) []TemplateName {
	files := make(map[TemplateName]interface{})

	for _, word := range strings.Fields(input) {
		if templateFile := GetTemplate(Input(word)); templateFile != TemplateName("") {
			fmt.Println("Added", templateFile)
			files[templateFile] = nil
		} else {
			fmt.Println("Could not find", word)
		}
	}

	fileSlice := []TemplateName{}

	for filename := range files {
		fileSlice = append(fileSlice, filename)
	}

	return fileSlice
}
