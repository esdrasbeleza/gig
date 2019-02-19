package main

import "strings"

type (
	Input        string
	TemplateFile string
)

func (i Input) Lowercase() Input {
	return Input(strings.ToLower(string(i)))
}

var (
	templateMap map[Input]TemplateFile
)

func init() {
	generateLanguageMaps()
}

func generateLanguageMaps() {
	inputAliasesMap := map[TemplateFile][]Input{
		"Go": {"Go", "golang"},
	}

	templateMap = make(map[Input]TemplateFile)

	for templateName, templateAliases := range inputAliasesMap {
		for _, templateAlias := range templateAliases {
			templateMap[templateAlias.Lowercase()] = templateName
		}
	}
}

func FileForLanguage(input Input) TemplateFile {
	if languageFile, exists := templateMap[input.Lowercase()]; exists {
		return languageFile
	}

	return TemplateFile("")
}
