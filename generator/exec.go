package generator

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCmdStream(cwd string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = cwd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command (%s %v) in %s: %w", name, args, cwd, err)
	}
	return nil
}

func YarnInstall(cwd string) error {
	// sama seperti pkg-install prefer yarn. Kita langsung panggil yarn.
	return RunCmdStream(cwd, "yarn", "install")
}
