package main

import "github.com/c-bata/go-prompt"

func completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{}

	for file, aliases := range inputAliasesMap {
		for _, alias := range aliases {
			suggest := prompt.Suggest{
				Text:        string(alias),
				Description: string(file),
			}

			suggestions = append(suggestions, suggest)
		}
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
