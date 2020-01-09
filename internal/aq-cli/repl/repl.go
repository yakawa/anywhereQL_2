package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yakawa1128/anywhereQL/pkg/lexer"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	r := bufio.NewScanner(in)
	fmt.Println("Start")
	for {
		fmt.Printf(PROMPT)
		lines := ""
		sf := r.Scan()
		if sf == false {
			return
		}
		for sf {
			line := r.Text()
			if line == "" {
				break
			}
			lines += "\n"
			lines += line
			sf = r.Scan()
			if sf == false {
				return
			}
		}
		if len(lines) != 0 {
			fmt.Println(lines)
			l := lexer.New(lines)
			tokens := l.Tokenize()
			for n, t := range tokens {
				fmt.Printf("[%d]: %s\n", n, t.Debug())
			}
		}
	}
}
