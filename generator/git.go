package generator

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CreateGitignoreFromTemplate(opts Options) error {
	// baca .gitignore dari template (embedded / override env)
	b, err := ReadTemplate(opts.Template, ".gitignore")
	if err != nil {
		return fmt.Errorf("read .gitignore template %s/.gitignore: %w", opts.Template, err)
	}

	dst := filepath.Join(opts.TargetDirectory, ".gitignore")
	if err := os.WriteFile(dst, b, 0o644); err != nil {
		return fmt.Errorf("write .gitignore %s: %w", dst, err)
	}
	return nil
}

func InitGit(opts Options) error {
	// Pastikan git tersedia
	if _, err := exec.LookPath("git"); err != nil {
		return errors.New("git not found in PATH")
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = opts.TargetDirectory

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = err.Error()
		}
		return fmt.Errorf("failed to initialize git: %s", msg)
	}

	return nil
}
