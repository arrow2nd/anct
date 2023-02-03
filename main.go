package main

import (
	"fmt"
	"os"

	"github.com/arrow2nd/anct/cmd"
	"github.com/arrow2nd/anct/config"
)

type exitCode int

const (
	exitCodeOK exitCode = iota
	exitCodeErrLoad
	exitCodeErrExec
)

func main() {
	cred, err := config.Load()
	if err != nil {
		exitError(err, int(exitCodeErrLoad))
	}

	c := cmd.New(cred)
	if err := c.Execute(); err != nil {
		exitError(err, int(exitCodeErrExec))
	}

	os.Exit(int(exitCodeOK))
}

func exitError(e error, c int) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", e.Error())
	os.Exit(c)
}
