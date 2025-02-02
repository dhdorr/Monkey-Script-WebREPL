package repl

import (
	"fmt"
	"io"

	"go-interpreter/evaluator"
	"go-interpreter/lexer"
	"go-interpreter/object"
	"go-interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, test_map *object.Environment, session_id, code string) string {
	//scanner := bufio.NewScanner(in)
	// env := object.NewEnvironment()

	//for {
	//fmt.Fprint(out, PROMPT)
	//scanned := scanner.Scan()
	//if !scanned {
	//	return
	//}

	//line := scanner.Text()
	// line := "2 + 2"
	l := lexer.New(code)

	// print lexical tokens
	// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
	// 	fmt.Fprintf(out, "%+v\n", tok)
	// }

	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		new_str := printParserErrors(out, p.Errors())
		//continue
		return new_str
	}

	evaluated := evaluator.Eval(program, test_map)
	result := "\n"
	if evaluated != nil {
		result = evaluated.Inspect()
		fmt.Printf(">> %s\n", result)
		// io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}

	// io.WriteString(out, program.String())
	// io.WriteString(out, "\n")
	//}
	return result
}

func printParserErrors(out io.Writer, errors []string) string {
	new_str := "Woops! We ran into some monkey business here!\n" + " parser errors:\n"
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")

	for _, msg := range errors {
		new_str += "\t" + msg + "\n"
		io.WriteString(out, "\t"+msg+"\n")
	}
	return new_str
}
