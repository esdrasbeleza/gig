package main

import "strings"

type (
	Input        string
	TemplateFile string
)

func (i Input) Lowercase() Input {
	return Input(strings.ToLower(string(i)))
}
