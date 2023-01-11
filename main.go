package main

import (
	"fmt"
	"os"

	"github.com/arrow2nd/annict-for-term/cmd"
)

const (
	ExitCodeOK int = iota
	ExitCodeErrHomeDir
	ExitCodeErrLoad
	ExitCodeErrExec
)

func main() {
	c := cmd.New()
	if err := c.Execute(); err != nil {
		printError(err)
		os.Exit(ExitCodeErrExec)
	}
}

func printError(e error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", e.Error())
}
