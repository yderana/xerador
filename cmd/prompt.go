package cmd

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/yderana/xerador/generator"
	"github.com/yderana/xerador/helpers"
)

func init() {
	promptui.IconGood = ""
	// promptui.IconBad = ""
	// promptui.IconSelect = ""
	// promptui.IconInitial = ""
}

// installChanged: true kalau user memang set --install / -i di CLI (biar gak ditanya lagi)
func PromptForMissingOptions(opts generator.Options, installChanged bool) (generator.Options, error) {
	// Kalau user belum pilih template via flag, kita tanya
	if opts.Template == "" {
		_, result, err := helpers.SelectPrompt("Please choose which you want", []string{"init repo", "create project", "create module"})
		if err != nil {
			return opts, err
		}

		// helpers.ClearPrevLines(1)

		switch result {
		case "init repo":
			opts.Template = "repo"
		case "create module":
			opts.Template = "module"
		default:
			opts.Template = "project"
		}

		// fmt.Println(gray(result))
	}

	// repo name (hanya untuk repo)
	if opts.Template == "repo" && opts.Repo == "" {
		v, err := helpers.TextPrompt("Repo Name", "repo-name")
		if err != nil {
			return opts, err
		}
		opts.Repo = v
	}

	// ProjectName:
	// - untuk "project" wajib
	// - untuk "repo" juga wajib (karena repoGenerator butuh projectName)
	// - untuk "module" juga wajib (karena moduleGenerator butuh projectName untuk cari targetService)
	if (opts.Template == "project" || opts.Template == "repo" || opts.Template == "module") && opts.ProjectName == "" {
		v, err := helpers.TextPrompt("Project Name", "project-name")
		if err != nil {
			return opts, err
		}
		opts.ProjectName = v
	}

	// ModuleName:
	// - untuk "module" wajib
	// - untuk "repo" juga wajib (repoGenerator bikin module default)
	// - untuk "project" juga wajib (serviceGenerator bikin module default)
	if (opts.Template == "module" || opts.Template == "repo" || opts.Template == "project") && opts.ModuleName == "" {
		v, err := helpers.TextPrompt("Module Name", "module-name")
		if err != nil {
			return opts, err
		}
		opts.ModuleName = v
	}

	// install prompt: hanya untuk repo & project, dan hanya kalau user belum set flag -i/--install
	if (opts.Template == "project" || opts.Template == "repo") && !installChanged {

		_, result, err := helpers.SelectPrompt("Run install package ?", []string{"Yes", "No"})
		if err != nil {
			return opts, err
		}

		helpers.ClearPrevLines(1)
		// set boolean once and print the selection
		opts.RunInstall = (result == "Yes")
		if opts.RunInstall {
			fmt.Println(helpers.Cyan("Run install package : ") + result)
		} else {
			fmt.Println(helpers.Gray("Run install package : " + result))
		}
	}

	// basic validation minimal
	switch opts.Template {
	case "repo":
		if opts.Repo == "" {
			return opts, errors.New("repo name is required")
		}
		if opts.ProjectName == "" {
			return opts, errors.New("project name is required")
		}
		if opts.ModuleName == "" {
			return opts, errors.New("module name is required")
		}
	case "project":
		if opts.ProjectName == "" {
			return opts, errors.New("project name is required")
		}
		if opts.ModuleName == "" {
			return opts, errors.New("module name is required")
		}
	case "module":
		if opts.ProjectName == "" {
			return opts, errors.New("project name is required")
		}
		if opts.ModuleName == "" {
			return opts, errors.New("module name is required")
		}
	default:
		return opts, errors.New("unknown template type")
	}

	return opts, nil
}
