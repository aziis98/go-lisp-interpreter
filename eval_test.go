package lisp_test

import (
	"fmt"
	"log"

	"github.com/aziis98/go-lisp-interpreter"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func ExampleInterpreter_Eval_string() {
	interp := lisp.NewInterpreter()
	v, err := interp.Eval(`"example"`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lisp.Show(v))

	// Output:
	// "example"
}

func ExampleInterpreter_Eval_quote1() {
	interp := lisp.NewInterpreter()
	v, err := interp.Eval(`#"example"`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lisp.Show(v))

	// Output:
	// #"example"
}

func ExampleInterpreter_Eval_quote2() {
	interp := lisp.NewInterpreter()
	v, err := interp.Eval(`#(+ 1 2 3)`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lisp.Show(v))

	// Output:
	// #(+ 1 2 3)
}

func ExampleInterpreter_Eval_plus() {
	interp := lisp.NewInterpreter()
	v, err := interp.Eval(`(+ 1 2 3)`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lisp.Show(v))

	// Output:
	// 6
}

func ExampleInterpreter_Eval_println() {
	interp := lisp.NewInterpreter()
	v, err := interp.Eval(`(fmt.Println "Hello, World!")`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lisp.Show(v))

	// Output:
	// Hello, World!
	// (14 <nil>)
}
