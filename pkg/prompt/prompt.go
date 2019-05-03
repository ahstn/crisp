package prompt

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ahstn/crisp/pkg/battery"
	"github.com/ahstn/crisp/pkg/git"
	"github.com/ahstn/crisp/pkg/kubernetes"
	"github.com/ahstn/crisp/pkg/path"
	"github.com/talal/go-bits/color"
)

// Info returns the prompt info line.
func Info() string {
	var line []string

	cwd, err := os.Getwd()
	if err != nil {
		cwd = os.Getenv("PWD")
	}
	cwd = filepath.Clean(cwd)

	line = appendUnlessEmpty(line, path.CurrentDirectory(cwd))
	line = appendUnlessEmpty(line, git.Branch(cwd))
	line = appendUnlessEmpty(line, kubernetes.Context())

	return strings.Join(line, " ")
}

// Symbol returns the prompt character
func Symbol() string {
	symbol := "❯"
	if overwrite := os.Getenv("CRISP_SYMBOL"); overwrite != "" {
		symbol = overwrite
	}

	return color.Sprintf(color.Magenta, "%v%v", battery.Percentage(), symbol)
}

func appendUnlessEmpty(list []string, val string) []string {
	if val == "" {
		return list
	}

	return append(list, val)
}
