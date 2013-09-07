package base

import (
	"os"
	"path/filepath"
)

func DataDir() string {
	return filepath.Join(os.Args[0], "..", "..")
}
