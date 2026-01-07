package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yderana/xerador/helpers"
)

func CreateProviders(opts Options) error {
	providersDir := filepath.Join(opts.TargetDirectory, "providers")
	if err := os.MkdirAll(providersDir, 0o755); err != nil {
		return fmt.Errorf("mkdir providers: %w", err)
	}

	title := helpers.UpperFirst(helpers.CamelCase(opts.ModuleName)) // Template
	variable := helpers.CamelCase(opts.ModuleName)                  // variable
	kebab := helpers.KebabCase(opts.ModuleName)                     // template

	relPath := filepath.Join("providers", "template.service.ts")
	b, err := ReadTemplate(opts.Template, relPath)
	if err != nil {
		return fmt.Errorf("read template %s/%s: %w", opts.Template, relPath, err)
	}

	s := string(b)
	s = strings.ReplaceAll(s, "Template", title)
	s = strings.ReplaceAll(s, "variable", variable)
	s = strings.ReplaceAll(s, "template", kebab)

	outPath := filepath.Join(providersDir, fmt.Sprintf("%s.service.ts", kebab))
	if err := os.WriteFile(outPath, []byte(s), 0o644); err != nil {
		return fmt.Errorf("write file %s: %w", outPath, err)
	}

	helpers.Sleep(1000)
	return nil
}
