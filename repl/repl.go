package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/iashyam/bclang/lexer"
	"github.com/iashyam/bclang/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if strings.Trim(line, " ") == "exit" {
			return
		}
		l := lexer.New(line)
		for true {
			tok := l.NextToken()
			fmt.Printf("%+v\n", tok)
			if tok.Type == token.EOF {
				break
			}
		}
	}
}
