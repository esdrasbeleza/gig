package main

var (
	templateMap map[Input]TemplateFile
)

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

func GetTemplate(input Input) TemplateFile {
	if languageFile, exists := templateMap[input.Lowercase()]; exists {
		return languageFile
	}

	return TemplateFile("")
}
