package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yderana/xerador/helpers"
)

func CreateInterfaces(opts Options) error {
	interfacesDir := filepath.Join(opts.TargetDirectory, "interfaces")
	if err := os.MkdirAll(interfacesDir, 0o755); err != nil {
		return fmt.Errorf("mkdir interfaces: %w", err)
	}

	title := helpers.UpperFirst(helpers.CamelCase(opts.ModuleName)) // Template
	kebab := helpers.KebabCase(opts.ModuleName)                     // template

	replaceTokens := func(s string) string {
		s = strings.ReplaceAll(s, "Template", title)
		s = strings.ReplaceAll(s, "template", kebab)
		return s
	}

	process := func(inRel, outName string) error {
		relPath := filepath.Join("interfaces", inRel)
		b, err := ReadTemplate(opts.Template, relPath)
		if err != nil {
			return fmt.Errorf("read template %s/%s: %w", opts.Template, relPath, err)
		}

		outPath := filepath.Join(interfacesDir, outName)
		if err := os.WriteFile(outPath, []byte(replaceTokens(string(b))), 0o644); err != nil {
			return fmt.Errorf("write file %s: %w", outPath, err)
		}
		return nil
	}

	// iparam-list
	if err := process(
		"iparam-list-template.ts",
		fmt.Sprintf("iparam-list-%s.ts", kebab),
	); err != nil {
		return err
	}

	// iparam
	if err := process(
		"iparam-template.ts",
		fmt.Sprintf("iparam-%s.ts", kebab),
	); err != nil {
		return err
	}

	helpers.Sleep(1000)
	return nil
}
