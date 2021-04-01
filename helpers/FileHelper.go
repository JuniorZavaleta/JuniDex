package helpers

import (
	"os"
	"path/filepath"
)

func GetTemplateFilepath(dir string, filename string) string {
	return filepath.Join(os.Getenv("DIR_TEMPLATE"), dir, filename)
}
