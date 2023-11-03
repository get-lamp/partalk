package main

import (
	"os"
	"partalk/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
