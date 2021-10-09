package main

import (
	"github.com/KaymeKaydex/lagrange-polynomial.git/command"
	"os"
)

func main() {
	os.Exit(command.Run(os.Args[1:]))
}
