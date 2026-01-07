package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yderana/xerador/helpers"
)

func CreateControllers(opts Options) error {
	// Pastikan folder target ada
	controllersDir := filepath.Join(opts.TargetDirectory, "controllers")
	if err := os.MkdirAll(controllersDir, 0o755); err != nil {
		return fmt.Errorf("mkdir controllers: %w", err)
	}

	// Hitung nama yang dibutuhkan untuk replace
	title := helpers.UpperFirst(helpers.CamelCase(opts.ModuleName)) // Template
	variable := helpers.CamelCase(opts.ModuleName)                  // variable
	kebab := helpers.KebabCase(opts.ModuleName)                     // template

	// Helper untuk proses 1 file template -> output
	process := func(inRel, outName string) error {
		relPath := filepath.Join("controllers", inRel)
		b, err := ReadTemplate(opts.Template, relPath)
		if err != nil {
			return fmt.Errorf("read template %s/%s: %w", opts.Template, relPath, err)
		}

		s := string(b)
		s = strings.ReplaceAll(s, "Template", title)
		s = strings.ReplaceAll(s, "variable", variable)
		s = strings.ReplaceAll(s, "template", kebab)

		outPath := filepath.Join(controllersDir, outName)
		if err := os.WriteFile(outPath, []byte(s), 0o644); err != nil {
			return fmt.Errorf("write file %s: %w", outPath, err)
		}
		return nil
	}

	// 1) controller.ts
	if err := process("template.controller.ts", fmt.Sprintf("%s.controller.ts", kebab)); err != nil {
		return err
	}

	// 2) resolvers.ts
	if err := process("template.resolvers.ts", fmt.Sprintf("%s.resolvers.ts", kebab)); err != nil {
		return err
	}

	// Samakan behavior Node: await sleep(1000)
	helpers.Sleep(1000)
	_ = time.Second // (optional) hapus kalau tidak dipakai; di sini Sleep sudah cukup

	return nil
}
