package main

import "github.com/c-bata/go-prompt"

func completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{}

	for key, template := range templateMap {
		mainSuggest := prompt.Suggest{
			Text:        string(key),
			Description: string(template.Name),
		}

		suggestions = append(suggestions, mainSuggest)
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
