package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yderana/xerador/helpers"
)

func CreateDirectory(opts Options) error {
	dirs := []string{
		opts.TargetDirectory,
		filepath.Join(opts.TargetDirectory, "interfaces"),
		filepath.Join(opts.TargetDirectory, "repository"),
		filepath.Join(opts.TargetDirectory, "providers"),
		filepath.Join(opts.TargetDirectory, "api", "graphql"),
		filepath.Join(opts.TargetDirectory, "api", "jsonschema"),
		filepath.Join(opts.TargetDirectory, "controllers"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("mkdir %s: %w", dir, err)
		}
	}

	helpers.Sleep(1000)
	return nil
}
