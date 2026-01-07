package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// ReadTemplate membaca file template berdasarkan template name:
// - "module" / "project" / "repo"
// relPath contoh:
// - "controllers/template.controller.ts"
// - "api/graphql/template.graphql"
//
// Urutan sumber:
//  1. Override folder via env XERADOR_TEMPLATES_DIR (opsional)
//     struktur: $XERADOR_TEMPLATES_DIR/<template>/<relPath>
//  2. Embedded templates (default & permanent)
func ReadTemplate(template string, relPath string) ([]byte, error) {
	// 1) override via env (optional)
	if base := os.Getenv("XERADOR_TEMPLATES_DIR"); base != "" {
		p := filepath.Join(base, template, relPath)
		if _, err := os.Stat(p); err == nil {
			return os.ReadFile(p)
		}
	}

	// 2) embedded
	full := filepath.Join("templates", template, relPath)
	b, err := templatesFS.ReadFile(full)
	if err != nil {
		return nil, fmt.Errorf("template not found: %s", full)
	}
	return b, nil
}

func CopyEmbeddedTemplateDir(template string, destDir string, clobber bool) error {
	root := filepath.ToSlash(filepath.Join("templates", template)) // embed pakai slash

	return fs.WalkDir(templatesFS, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// hitung relative path dari templates/<template>/
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}

		dstPath := filepath.Join(destDir, filepath.FromSlash(rel))

		if d.IsDir() {
			return os.MkdirAll(dstPath, 0o755)
		}

		// pastikan folder parent ada
		if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
			return err
		}

		// jika clobber=false dan file sudah ada, skip
		if !clobber {
			if _, err := os.Stat(dstPath); err == nil {
				return nil
			}
		}

		b, err := templatesFS.ReadFile(path)
		if err != nil {
			return err
		}

		// permission default 0644 (kalau kamu perlu executable, bisa ditambah rule khusus)
		if err := os.WriteFile(dstPath, b, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", dstPath, err)
		}

		return nil
	})
}
