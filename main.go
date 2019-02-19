package main

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
)

func main() {
	files := []TemplateName{}

	fmt.Println("Please select language or platform:")

	for {
		input := prompt.Input("> ", completer)

		if strings.TrimSpace(input) == "" {
			break
		}

		newFiles := ParseInput(input)

		files = append(files, newFiles...)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
