package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yderana/xerador/helpers"
)

func CreateModuleFiles(opts Options) error {
	title := helpers.UpperFirst(helpers.CamelCase(opts.ModuleName)) // Template
	kebab := helpers.KebabCase(opts.ModuleName)                     // template

	// 1) generate <module>.module.ts dari template.module.ts
	{
		// dulu: filepath.Join(opts.TemplateDirectory, "template.module.ts")
		b, err := ReadTemplate(opts.Template, "template.module.ts")
		if err != nil {
			return fmt.Errorf("read template %s/template.module.ts: %w", opts.Template, err)
		}

		s := string(b)
		s = strings.ReplaceAll(s, "Template", title)
		s = strings.ReplaceAll(s, "template", kebab)

		outPath := filepath.Join(opts.TargetDirectory, fmt.Sprintf("%s.module.ts", kebab))
		if err := os.WriteFile(outPath, []byte(s), 0o644); err != nil {
			return fmt.Errorf("write file %s: %w", outPath, err)
		}
	}

	// Kalau ini mode "module", update app.module.ts + package.json + tsconfig.json di targetService
	if opts.Template == "module" {
		appModulePath := filepath.Join(opts.TargetService, "src", "app.module.ts")

		// A) sisipkan import module sebelum import ScheduleModule
		{
			b, err := os.ReadFile(appModulePath)
			if err != nil {
				return fmt.Errorf("read %s: %w", appModulePath, err)
			}
			data := string(b)

			needle := "import { ScheduleModule } from '@nestjs/schedule';"
			parts := strings.Split(data, needle)
			if len(parts) >= 2 {
				insert := fmt.Sprintf(
					`import { %sModule } from "@%s/%s.module";`+"\n"+`%s`,
					title, kebab, kebab, needle,
				)
				newData := parts[0] + insert + strings.Join(parts[1:], needle)
				if err := os.WriteFile(appModulePath, []byte(newData), 0o644); err != nil {
					return fmt.Errorf("write %s: %w", appModulePath, err)
				}
			}
		}

		helpers.Sleep(1000)

		// B) sisipkan "<Module>Module," sebelum "ScheduleModule.forRoot()"
		{
			b, err := os.ReadFile(appModulePath)
			if err != nil {
				return fmt.Errorf("read %s: %w", appModulePath, err)
			}
			data := string(b)

			needle := "ScheduleModule.forRoot()"
			parts := strings.Split(data, needle)
			if len(parts) >= 2 {
				insert := fmt.Sprintf("%sModule,\n%s", title, needle)
				newData := parts[0] + insert + strings.Join(parts[1:], needle)
				if err := os.WriteFile(appModulePath, []byte(newData), 0o644); err != nil {
					return fmt.Errorf("write %s: %w", appModulePath, err)
				}
			}
		}

		// C) update package.json -> jest.moduleNameMapper tambah:
		// "@<module>/(.*)": "<rootDir>/<module>/$1"
		{
			pkgPath := filepath.Join(opts.TargetService, "package.json")
			b, err := os.ReadFile(pkgPath)
			if err != nil {
				return fmt.Errorf("read %s: %w", pkgPath, err)
			}

			var root map[string]any
			if err := json.Unmarshal(b, &root); err != nil {
				return fmt.Errorf("parse %s: %w", pkgPath, err)
			}

			jest, _ := root["jest"].(map[string]any)
			if jest == nil {
				jest = map[string]any{}
				root["jest"] = jest
			}

			mapper, _ := jest["moduleNameMapper"].(map[string]any)
			if mapper == nil {
				mapper = map[string]any{}
			}

			key := fmt.Sprintf("@%s/(.*)", kebab)
			val := fmt.Sprintf("<rootDir>/%s/$1", kebab)
			mapper[key] = val
			jest["moduleNameMapper"] = mapper

			out, err := json.Marshal(root)
			if err != nil {
				return fmt.Errorf("marshal %s: %w", pkgPath, err)
			}
			if err := os.WriteFile(pkgPath, out, 0o644); err != nil {
				return fmt.Errorf("write %s: %w", pkgPath, err)
			}
		}

		// D) update tsconfig.json -> compilerOptions.paths tambah:
		// "@<module>/*": ["<module>/*"]
		{
			tsPath := filepath.Join(opts.TargetService, "tsconfig.json")
			b, err := os.ReadFile(tsPath)
			if err != nil {
				return fmt.Errorf("read %s: %w", tsPath, err)
			}

			var root map[string]any
			if err := json.Unmarshal(b, &root); err != nil {
				return fmt.Errorf("parse %s: %w", tsPath, err)
			}

			co, _ := root["compilerOptions"].(map[string]any)
			if co == nil {
				co = map[string]any{}
				root["compilerOptions"] = co
			}

			paths, _ := co["paths"].(map[string]any)
			if paths == nil {
				paths = map[string]any{}
			}

			key := fmt.Sprintf("@%s/*", kebab)
			paths[key] = []any{fmt.Sprintf("%s/*", kebab)}
			co["paths"] = paths

			out, err := json.Marshal(root)
			if err != nil {
				return fmt.Errorf("marshal %s: %w", tsPath, err)
			}
			if err := os.WriteFile(tsPath, out, 0o644); err != nil {
				return fmt.Errorf("write %s: %w", tsPath, err)
			}
		}
	}

	helpers.Sleep(1000)
	return nil
}

func ImplementModule(opts Options) error {
	kebabProject := helpers.KebabCase(opts.ProjectName)
	kebabModule := helpers.KebabCase(opts.ModuleName)

	// Pastikan folder yang dibutuhkan ada
	if err := os.MkdirAll(filepath.Join(opts.TargetDirectory, "src"), 0o755); err != nil {
		return fmt.Errorf("mkdir src: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(opts.TargetDirectory, "documentation"), 0o755); err != nil {
		return fmt.Errorf("mkdir documentation: %w", err)
	}

	// package.json
	{
		b, err := ReadTemplate(opts.Template, "package.json")
		if err != nil {
			return fmt.Errorf("read template %s/package.json: %w", opts.Template, err)
		}
		s := string(b)
		s = strings.ReplaceAll(s, "project-template", kebabProject)
		s = strings.ReplaceAll(s, "template", kebabModule)

		outPath := filepath.Join(opts.TargetDirectory, "package.json")
		if err := os.WriteFile(outPath, []byte(s), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", outPath, err)
		}
	}

	// tsconfig.json
	{
		b, err := ReadTemplate(opts.Template, "tsconfig.json")
		if err != nil {
			return fmt.Errorf("read template %s/tsconfig.json: %w", opts.Template, err)
		}
		s := string(b)
		s = strings.ReplaceAll(s, "project-template", kebabProject)
		s = strings.ReplaceAll(s, "template", kebabModule)

		outPath := filepath.Join(opts.TargetDirectory, "tsconfig.json")
		if err := os.WriteFile(outPath, []byte(s), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", outPath, err)
		}
	}

	// README.md
	{
		b, err := ReadTemplate(opts.Template, "README.md")
		if err != nil {
			return fmt.Errorf("read template %s/README.md: %w", opts.Template, err)
		}
		s := string(b)
		s = strings.ReplaceAll(s, "project-template", helpers.StartCase(opts.ProjectName))
		s = strings.ReplaceAll(s, "template", kebabModule)

		outPath := filepath.Join(opts.TargetDirectory, "README.md")
		if err := os.WriteFile(outPath, []byte(s), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", outPath, err)
		}
	}

	// src/app.module.ts
	{
		rel := filepath.Join("src", "app.module.ts")
		b, err := ReadTemplate(opts.Template, rel)
		if err != nil {
			return fmt.Errorf("read template %s/%s: %w", opts.Template, rel, err)
		}

		title := helpers.UpperFirst(helpers.CamelCase(opts.ModuleName))
		s := string(b)
		s = strings.ReplaceAll(s, "Template", title)
		s = strings.ReplaceAll(s, "template", kebabModule)

		outPath := filepath.Join(opts.TargetDirectory, "src", "app.module.ts")
		if err := os.WriteFile(outPath, []byte(s), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", outPath, err)
		}
	}

	// documentation/Environment.postman_environment.json
	{
		rel := filepath.Join("documentation", "Environment.postman_environment.json")
		b, err := ReadTemplate(opts.Template, rel)
		if err != nil {
			return fmt.Errorf("read template %s/%s: %w", opts.Template, rel, err)
		}

		s := string(b)
		s = strings.ReplaceAll(s, "project-template", kebabProject)
		// Node: replace "template" dengan kebabCase(projectName)
		s = strings.ReplaceAll(s, "template", kebabProject)

		outPath := filepath.Join(opts.TargetDirectory, "documentation", "Environment.postman_environment.json")
		if err := os.WriteFile(outPath, []byte(s), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", outPath, err)
		}
	}

	helpers.Sleep(1000)
	return nil
}
