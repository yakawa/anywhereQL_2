package main

import (
	"os"

	"github.com/yakawa1128/anywhereQL/internal/aq-cli/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
