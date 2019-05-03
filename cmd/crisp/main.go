package main

import (
	"os"

	"github.com/ahstn/crisp/pkg/prompt"
)

func main() {
	os.Stdout.Write([]byte("\n" + prompt.Info() + "\n" + prompt.Symbol()))
}
