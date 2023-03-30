package main

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
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

	fs.WalkDir(templates, ".", func(fullPath string, file fs.DirEntry, err error) error {
		info, _ := file.Info()

		if info.IsDir() {
			return nil
		}

		var (
			extension        = filepath.Ext(file.Name())
			withoutExtension = strings.TrimSuffix(file.Name(), extension)
			beforePlus       = strings.Split(withoutExtension, "+")[0]
			templateName     = Key(beforePlus)
		)

		if currentTemplate, exists := templateMap[templateName.Lowercase()]; exists {
			currentTemplate.Files = append(currentTemplate.Files, fullPath)
		} else {
			templateMap[templateName.Lowercase()] = &Template{
				Name:  templateName,
				Files: []string{fullPath},
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
