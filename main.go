package main

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
)

func main() {
	files := NewKeySet()

	fmt.Println("Please select language or platform:")

	for {
		prefix := fmt.Sprintf("[%s] > ", strings.Join(files.Strings(), ", "))
		input := prompt.Input(prefix, completer)

		if strings.TrimSpace(input) == "" {
			break
		}

		newFiles := ParseInput(input)

		files.Add(newFiles...)
	}

	// TODO: create file
}
