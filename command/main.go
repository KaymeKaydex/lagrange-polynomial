package command

import (
	"io"
	"os"
)

type RunOptions struct {
	Stdout io.Writer
	Stderr io.Writer
}

func Run(args []string) int {
	return RunCustom(args, nil)
}

func RunCustom(args []string, runOpts *RunOptions) int {
	if runOpts == nil {
		runOpts = &RunOptions{}
	}
	if runOpts.Stdout == nil {
		runOpts.Stdout = os.Stdout
	}
	if runOpts.Stderr == nil {
		runOpts.Stderr = os.Stderr
	}

	return 0
}
