package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yderana/xerador/helpers"
)

func CreateRepository(opts Options) error {
	repoDir := filepath.Join(opts.TargetDirectory, "repository")
	if err := os.MkdirAll(repoDir, 0o755); err != nil {
		return fmt.Errorf("mkdir repository: %w", err)
	}

	title := helpers.UpperFirst(helpers.CamelCase(opts.ModuleName)) // Template
	kebab := helpers.KebabCase(opts.ModuleName)                     // template

	replaceTokens := func(s string) string {
		s = strings.ReplaceAll(s, "Template", title)
		s = strings.ReplaceAll(s, "template", kebab)
		return s
	}

	process := func(inRel, outName string) error {
		relPath := filepath.Join("repository", inRel)

		b, err := ReadTemplate(opts.Template, relPath)
		if err != nil {
			return fmt.Errorf("read template %s/%s: %w", opts.Template, relPath, err)
		}

		outPath := filepath.Join(repoDir, outName)
		if err := os.WriteFile(outPath, []byte(replaceTokens(string(b))), 0o644); err != nil {
			return fmt.Errorf("write file %s: %w", outPath, err)
		}
		return nil
	}

	// 1) entity
	if err := process(
		"template.entity.ts",
		fmt.Sprintf("%s.entity.ts", kebab),
	); err != nil {
		return err
	}

	// 2) repository
	if err := process(
		"template.repository.ts",
		fmt.Sprintf("%s.repository.ts", kebab),
	); err != nil {
		return err
	}

	helpers.Sleep(1000)
	return nil
}
