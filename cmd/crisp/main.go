package main

import (
	"os"

	"github.com/ahstn/crisp/pkg/prompt"
	"github.com/talal/go-bits/color"
)

func main() {
	if len(os.Args) > 1 {
		color.Fprintln(os.Stderr, color.Red, "Prompt error: no arguing with Crisp")
		os.Exit(1)
	}

	os.Stdout.Write([]byte("\n" + prompt.Info() + "\n" + prompt.Symbol()))
}
