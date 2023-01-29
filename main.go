package main

import (
	"fmt"
	"os"

	"github.com/arrow2nd/anct/cmd"
	"github.com/arrow2nd/anct/credencial"
)

const (
	ExitCodeOK int = iota
	ExitCodeErrLoad
	ExitCodeErrExec
)

func main() {
	cred, err := credencial.Load()
	if err != nil {
		exitError(err, ExitCodeErrLoad)
	}

	c := cmd.New(cred)
	if err := c.Execute(); err != nil {
		exitError(err, ExitCodeErrExec)
	}
}

func exitError(e error, c int) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", e.Error())
	os.Exit(c)
}
