package ns

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gizahNL/gojail"
)

const (
	defaultJailDir = "/var/lib/runj/jails/"
	defaultLinkName = "netns"
	defaultNsDir = "/var/run/netns/"
)

func GetCurrentNsJail(childJail gojail.Jail) (gojail.Jail, error) {
	nsName, err := getNsPath(childJail.Name)
	if err != nil {
		return nil, err
	}

	nsJail, err := gojail.JailGetByName(nsName)
	if err != nil {
		return nil, err
	}

	return nsJail, nil
}

func getNsPath(jailID string) (string, error) {
	path := filepath.Join(defaultJailDir, jailID, defaultLinkName)

	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	var realPath string
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		realPath, err = os.Readlink(path)
		if err != nil {
			return "", err
		}
	} else{
		return "", fmt.Errorf("Path is not symboliclink.")
	}
	return filepath.Base(realPath), nil
}
