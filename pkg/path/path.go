package path

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/talal/go-bits/color"
)

// CurrentDirectory returns the current working directory.
func CurrentDirectory(cwd string) string {
	return color.Sprintf(color.Blue, sanitiseDirectory(cwd))
}

func sanitiseDirectory(cwd string) string {
	if cwd == "/" {
		return cwd
	}

	if info, err := os.Stat(cwd); !info.IsDir() && err == nil {
		cwd = filepath.Dir(cwd)
	} else if err != nil {
		return cwd
	}

	return strings.Replace(cwd, os.Getenv("HOME"), "~", 1)
}
