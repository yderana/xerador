package generator

import (
	"fmt"
	"os"

	"github.com/yderana/xerador/helpers"
)

func CreateService(opts Options) error {
	// pastikan target directory ada
	if err := os.MkdirAll(opts.TargetDirectory, 0o755); err != nil {
		return fmt.Errorf("mkdir targetDirectory: %w", err)
	}

	helpers.Sleep(1000)

	// copy templates/<template> -> targetDirectory (clobber = false)
	// biasanya opts.Template == "project"
	if err := CopyEmbeddedTemplateDir(opts.Template, opts.TargetDirectory, false); err != nil {
		return fmt.Errorf("copy service template: %w", err)
	}

	helpers.Sleep(2000)
	return nil
}
