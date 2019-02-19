package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type TemplateFile string

type Template struct {
	Name    TemplateName
	Aliases []Input
	Files   []string
}

var (
	templateFileMap map[TemplateName]*Template

	inputAliasesMap = map[TemplateName][]Input{
		"Go":               {"Go", "golang"},
		"VisualStudioCode": {"VisualStudioCode", "VSCode", "code"},
	}

	templateMap map[Input]*Template
)

func generateTemplateFileMap() {
	templateFileMap = make(map[TemplateName]*Template)

	files, err := ioutil.ReadDir("./gitignore/templates")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		var (
			extension        = path.Ext(file.Name())
			withoutExtension = strings.TrimSuffix(file.Name(), extension)
			beforePlus       = strings.Split(withoutExtension, "+")[0]
			templateName     = TemplateName(beforePlus)
			newFile          = file.Name()
		)

		if currentTemplate, exists := templateFileMap[templateName]; exists {
			currentTemplate.Files = append(currentTemplate.Files, newFile)
		} else {
			templateFileMap[templateName] = &Template{
				Name:  templateName,
				Files: []string{newFile},
			}
		}
	}
}

func generateLanguageMaps() {
	templateMap = make(map[Input]*Template)

	for templateName, template := range templateFileMap {
		if templateAliases, exists := inputAliasesMap[templateName]; exists {
			template.Aliases = templateAliases

			for _, templateAlias := range templateAliases {
				templateMap[templateAlias.Lowercase()] = template
			}
		}
	}
}

func GetTemplate(input Input) *Template {
	if template, exists := templateMap[input.Lowercase()]; exists {
		return template
	}

	return nil
}

func ParseInput(input string) []TemplateName {
	files := make(map[TemplateName]interface{})

	for _, word := range strings.Fields(input) {
		if template := GetTemplate(Input(word)); template != nil {
			fmt.Println("Added", template.Name)
			files[template.Name] = nil
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
