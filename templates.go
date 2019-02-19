package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Template struct {
	Name  Key
	Files []string
}

var (
	templateMap      map[Key]*Template
	templateAliasMap map[Key]*Template

	aliases = map[Key][]Key{
		"Go":               {"golang"},
		"VisualStudioCode": {"VSCode", "code"},
	}
)

func generateTemplateFileMap() {
	templateMap = make(map[Key]*Template)

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
			templateName     = Key(beforePlus)
			newFile          = file.Name()
		)

		if currentTemplate, exists := templateMap[templateName]; exists {
			currentTemplate.Files = append(currentTemplate.Files, newFile)
		} else {
			templateMap[templateName] = &Template{
				Name:  templateName,
				Files: []string{newFile},
			}
		}
	}
}

func addAliases() {
	templateAliasMap = make(map[Key]*Template)

	for templateName, template := range templateMap {
		templateAliasMap[templateName.Lowercase()] = template

		if templateAliases, exists := aliases[templateName]; exists {
			for _, templateAlias := range templateAliases {
				templateAliasMap[templateAlias] = template
				templateAliasMap[templateAlias.Lowercase()] = template
			}
		}
	}
}

func GetTemplate(input Key) *Template {
	if template, exists := templateMap[Key(string(input))]; exists {
		return template
	}

	if template, exists := templateAliasMap[Key(string(input))]; exists {
		return template
	}

	return nil
}

func ParseInput(input string) []Key {
	files := make(map[Key]interface{})

	for _, word := range strings.Fields(input) {
		if template := GetTemplate(Key(word)); template != nil {
			files[template.Name] = nil
		} else {
			fmt.Println("Could not find", word)
		}
	}

	fileSlice := []Key{}

	for filename := range files {
		fileSlice = append(fileSlice, filename)
	}

	return fileSlice
}
