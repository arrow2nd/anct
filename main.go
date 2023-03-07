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
	exitCodeErrGetDir
	exitCodeErrLoad
	exitCodeErrExec
)

func main() {
	cfg, err := config.New()
	if err != nil {
		exitError(err, exitCodeErrGetDir)
	}

	c, err := cmd.New(cfg)
	if err != nil {
		exitError(err, exitCodeErrLoad)
	}

	if err := c.Execute(); err != nil {
		exitError(err, exitCodeErrExec)
	}

	os.Exit(int(exitCodeOK))
}

// exitError : エラーを出力して終了
func exitError(e error, c exitCode) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", e.Error())
	os.Exit(int(c))
}
