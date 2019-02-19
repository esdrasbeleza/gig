package main

import "strings"

type (
	Input        string
	TemplateName string
)

func (i Input) Lowercase() Input {
	return Input(strings.ToLower(string(i)))
}
