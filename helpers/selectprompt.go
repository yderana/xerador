package helpers

import "github.com/manifoldco/promptui"

func SelectPrompt(label string, items []string) (int, string, error) {
	p := promptui.Select{
		Label: label,
		Items: items,
		Size:  5,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "▸ {{ . | cyan }}",
			Inactive: "  {{ . | faint }}",
			Selected: "{{ . | green }}",
		},
	}
	return p.Run()
}
