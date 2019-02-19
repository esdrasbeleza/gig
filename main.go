package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func main() {
	var files KeySet

	if len(os.Args) > 1 {
		files = getKeysFromArgs()
	} else {
		files = getKeysFromPrompt()
	}

	for _, key := range files.Keys() {
		filesToOutput := GetTemplate(key).Files

		for _, file := range filesToOutput {
			content, _ := box.FindString(file)
			fmt.Println(content)
		}
	}
}

func getKeysFromArgs() KeySet {
	var (
		params = strings.Join(os.Args[1:], " ")
		files  = ParseInput(params)
	)

	return files
}

func getKeysFromPrompt() KeySet {
	files := NewKeySet()

	fmt.Println("Please select language or platform:")

	for {
		var (
			prefix = fmt.Sprintf("[%s] > ", strings.Join(files.Strings(), ", "))
			input  = prompt.Input(prefix, completer)
		)

		if strings.TrimSpace(input) == "" {
			break
		}

		newFiles := ParseInput(input)

		files.Merge(newFiles)
	}

	return files
}
