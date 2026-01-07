package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type LicenseOptions struct {
	Name  string
	Email string
}

// Jika mau, ini bisa digabung ke Options utama
// (aku pisahkan dulu biar jelas mapping-nya)
func CreateLicense(opts Options, lic LicenseOptions) error {
	targetPath := filepath.Join(opts.TargetDirectory, "LICENSE")

	year := time.Now().Year()
	holder := fmt.Sprintf("%s (%s)", lic.Name, lic.Email)

	content := mitLicenseTemplate()
	content = strings.ReplaceAll(content, "<year>", fmt.Sprintf("%d", year))
	content = strings.ReplaceAll(content, "<copyright holders>", holder)

	if err := os.WriteFile(targetPath, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write LICENSE: %w", err)
	}

	return nil
}

func mitLicenseTemplate() string {
	return `MIT License

Copyright (c) <year> <copyright holders>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`
}
