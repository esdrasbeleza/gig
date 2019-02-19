package main

import packr "github.com/gobuffalo/packr/v2"

func init() {
	loadTemplates()
	generateTemplateFileMap()
	addAliases()
}

var box *packr.Box

func loadTemplates() {
	box = packr.New("templates", "gitignore/templates")
}
