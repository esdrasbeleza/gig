package main

import (
	"fmt"
	"strings"
)

var (
	inputAliasesMap = map[TemplateFile][]Input{
		"Go":         {"Go", "golang"},
		"JavaScript": {"JavaScript", "JS"},
		"TypeScript": {"TypeScript", "TS"},
	}

	templateMap map[Input]TemplateFile
)

func generateLanguageMaps() {
	templateMap = make(map[Input]TemplateFile)

	for templateName, templateAliases := range inputAliasesMap {
		for _, templateAlias := range templateAliases {
			templateMap[templateAlias.Lowercase()] = templateName
		}
	}
}

func GetTemplate(input Input) TemplateFile {
	if languageFile, exists := templateMap[input.Lowercase()]; exists {
		return languageFile
	}

	return TemplateFile("")
}

func ParseInput(input string) []TemplateFile {
	files := make(map[TemplateFile]interface{})

	for _, word := range strings.Fields(input) {
		if templateFile := GetTemplate(Input(word)); templateFile != TemplateFile("") {
			fmt.Println("Added", templateFile)
			files[templateFile] = nil
		} else {
			fmt.Println("Could not find", word)
		}
	}

	fileSlice := []TemplateFile{}

	for filename := range files {
		fileSlice = append(fileSlice, filename)
	}

	return fileSlice
}
