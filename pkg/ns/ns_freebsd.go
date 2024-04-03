package ns

import (
	"fmt"
	"os"
	"path/filepath"
)

func JailGetByPath(nsPath string) (string, error) {
	info, err := os.Lstat(nsPath)
	if err != nil {
		return "", err
	}

	var link string
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		link, err = os.Readlink(nsPath)
		if err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("Path is not Symboliclink.")
	}

	return filepath.Base(link), nil
}
