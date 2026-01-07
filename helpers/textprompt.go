package helpers

import "github.com/manifoldco/promptui"

func TextPrompt(label, def string) (string, error) {
	p := promptui.Prompt{
		Label:   label, // tanpa "?"
		Default: def,
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . | cyan }} : ",
			Valid:   "{{ . | cyan }} : ",
			Invalid: "{{ . | cyan }} : ",
			Success: "{{ . | cyan }} : ",
		},
	}
	return p.Run()
}
