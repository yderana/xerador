package helpers

import "os"

func PathExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
