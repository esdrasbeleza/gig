package main

import (
	"embed"
)

func init() {
	generateTemplateFileMap()
	addAliases()
}

//go:embed gitignore/templates
var templates embed.FS
