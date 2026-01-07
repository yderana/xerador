package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yderana/xerador/helpers"
)

func CreateSchema(opts Options) error {
	graphqlDir := filepath.Join(opts.TargetDirectory, "api", "graphql")
	jsonschemaDir := filepath.Join(opts.TargetDirectory, "api", "jsonschema")

	if err := os.MkdirAll(graphqlDir, 0o755); err != nil {
		return fmt.Errorf("mkdir api/graphql: %w", err)
	}
	if err := os.MkdirAll(jsonschemaDir, 0o755); err != nil {
		return fmt.Errorf("mkdir api/jsonschema: %w", err)
	}

	title := helpers.UpperFirst(helpers.CamelCase(opts.ModuleName)) // Template
	variable := helpers.CamelCase(opts.ModuleName)                  // variable
	kebab := helpers.KebabCase(opts.ModuleName)                     // template

	replaceAll := func(s string, includeVariable bool) string {
		s = strings.ReplaceAll(s, "Template", title)
		if includeVariable {
			s = strings.ReplaceAll(s, "variable", variable)
		}
		s = strings.ReplaceAll(s, "template", kebab)
		return s
	}

	process := func(relPath string, outDir string, outName string, includeVariable bool) error {
		b, err := ReadTemplate(opts.Template, relPath)
		if err != nil {
			return fmt.Errorf("read template %s/%s: %w", opts.Template, relPath, err)
		}

		outPath := filepath.Join(outDir, outName)
		if err := os.WriteFile(outPath, []byte(replaceAll(string(b), includeVariable)), 0o644); err != nil {
			return fmt.Errorf("write file %s: %w", outPath, err)
		}
		return nil
	}

	// 1) GraphQL schema
	{
		rel := filepath.Join("api", "graphql", "template.graphql")
		if err := process(
			rel,
			graphqlDir,
			fmt.Sprintf("%s.graphql", kebab),
			true, // include variable
		); err != nil {
			return err
		}
	}

	// 2) JSON schema list
	{
		rel := filepath.Join("api", "jsonschema", "template.schema-list.ts")
		if err := process(
			rel,
			jsonschemaDir,
			fmt.Sprintf("%s.schema-list.ts", kebab),
			false,
		); err != nil {
			return err
		}
	}

	// 3) JSON schema request
	{
		rel := filepath.Join("api", "jsonschema", "template.schema-request.ts")
		if err := process(
			rel,
			jsonschemaDir,
			fmt.Sprintf("%s.schema-request.ts", kebab),
			false,
		); err != nil {
			return err
		}
	}

	helpers.Sleep(1000)
	return nil
}
