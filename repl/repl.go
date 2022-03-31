package repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"zabaan/evaluator"
	"zabaan/lexer"
	"zabaan/object"
	"zabaan/parser"
)

func Start(username string, in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf("\033[1;32mcommand me \033[35m%s$\033[0m ", username)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const LAAMS_LOGO = `
--- Laams LLC: Zabaan Programming Language ---

🅻 🅰 🅰 🅼 🆂
██╗░░░░░░█████╗░░█████╗░███╗░░░███╗░██████╗
██║░░░░░██╔══██╗██╔══██╗████╗░████║██╔════╝
██║░░░░░███████║███████║██╔████╔██║╚█████╗░
██║░░░░░██╔══██║██╔══██║██║╚██╔╝██║░╚═══██╗
███████╗██║░░██║██║░░██║██║░╚═╝░██║██████╔╝
╚══════╝╚═╝░░╚═╝╚═╝░░╚═╝╚═╝░░░░░╚═╝╚═════╝░

`

func printParserErrors(out io.Writer, errors []string) {
	ioutil.ReadFile("file.")
	io.WriteString(out, LAAMS_LOGO)
	io.WriteString(out, "Woops! Zabaan ran into some problems here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
