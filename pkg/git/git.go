package git

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/talal/go-bits/color"
)

const dirtySymbol = "*"

func Branch(cwd string) string {
	gitDir, _ := findGitRepo(cwd)
	fmt.Println("gitdir:" + gitDir)
	if gitDir == "" {
		return ""
	}

	bytes, err := ioutil.ReadFile(filepath.Join(gitDir, "HEAD"))
	if err != nil {
		fmt.Println(err)
		return "unknown"
	}
	refSpec := strings.TrimSpace(string(bytes))

	// detached HEAD?
	if !strings.HasPrefix(refSpec, "ref: refs/") {
		return "detached"
	}

	refSpecDisplay := strings.TrimPrefix(refSpec, "ref: refs/heads/")

	if ok, err := isDirty(); ok && err == nil {
		refSpecDisplay += dirtySymbol
	}
	return color.Sprintf(color.White, refSpecDisplay)
}

func isDirty() (bool, error) {
	var stdout bytes.Buffer
	cmd := exec.Command("git", "status", "--porcelain", "--ignore-submodules", "-unormal")
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return false, err
	}

	if string(stdout.Bytes()) != "" {
		return true, nil
	}

	return false, nil
}

func findGitRepo(path string) (string, error) {
	gitEntry := filepath.Join(path, ".git")
	fi, err := os.Stat(gitEntry)
	switch {
	case err == nil:
		//found - continue below with further checks
	case !os.IsNotExist(err):
		return "", err
	case path == "/":
		return "", nil
	default:
		return findGitRepo(filepath.Dir(path))
	}

	if !fi.IsDir() {
		return "", nil
	}

	return gitEntry, nil
}
