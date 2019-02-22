package main

import (
	"log"
	"path"
	"strings"

	"github.com/gobuffalo/packr/v2/file"
)

type Template struct {
	Name  Key
	Files []string
}

var (
	templateMap map[Key]*Template
	aliases     = map[Key][]Key{
		"Go":               {"golang"},
		"VisualStudioCode": {"VSCode", "code"},
	}
)

func generateTemplateFileMap() {
	templateMap = make(map[Key]*Template)

	box.Walk(func(filename string, file file.File) error {
		info, _ := file.FileInfo()

		if info.IsDir() {
			return nil
		}

		var (
			extension        = path.Ext(file.Name())
			withoutExtension = strings.TrimSuffix(file.Name(), extension)
			beforePlus       = strings.Split(withoutExtension, "+")[0]
			templateName     = Key(beforePlus)
			newFile          = file.Name()
		)

		if currentTemplate, exists := templateMap[templateName.Lowercase()]; exists {
			currentTemplate.Files = append(currentTemplate.Files, newFile)
		} else {
			templateMap[templateName.Lowercase()] = &Template{
				Name:  templateName,
				Files: []string{newFile},
			}
		}

		return nil
	})
}

func addAliases() {
	for _, template := range templateMap {
		if templateAliases, exists := aliases[template.Name]; exists {
			for _, templateAlias := range templateAliases {
				templateMap[templateAlias.Lowercase()] = template
			}
		}
	}
}

func GetTemplate(input Key) *Template {
	if template, exists := templateMap[input.Lowercase()]; exists {
		return template
	}

	return nil
}

func ParseInput(input string) KeySet {
	files := NewKeySet()

	for _, word := range strings.Fields(input) {
		if template := GetTemplate(Key(word)); template != nil {
			files.Add(template.Name)
		} else {
			log.Println("Could not find", word)
		}
	}

	return files
}
