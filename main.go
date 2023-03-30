package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func main() {
	if len(os.Args) > 1 {
		getKeysFromArgs()
	} else {
		getKeysFromPrompt()
	}
}

func getKeysFromArgs() {
	var (
		params = strings.Join(os.Args[1:], " ")
		files  = ParseInput(params)
	)

	for _, key := range files.Keys() {
		filesToOutput := GetTemplate(key).Files

		for _, file := range filesToOutput {
			content, err := templates.ReadFile(file)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			fmt.Println(string(content))
		}
	}
}

func getKeysFromPrompt() {
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

	if files.IsEmpty() {
		os.Exit(0)
	}

	var (
		filenameInput   = prompt.Input("Output file to append: ", nilCompleter)
		outputFile, err = os.OpenFile(filenameInput, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	defer outputFile.Close()

	for _, templateFileKeys := range files.Keys() {
		filesForKey := GetTemplate(templateFileKeys).Files

		for _, templateFile := range filesForKey {
			file, err := templates.Open(templateFile)
			if err != nil {
				log.Println(err.Error())
				os.Exit(1)
			}
			defer file.Close()

			if _, err := io.Copy(outputFile, file); err != nil {
				log.Println(err.Error())
				os.Exit(1)
			}
		}
	}

}
