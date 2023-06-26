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

func ExampleInterpreter_Execute_list() {
	interp := lisp.NewInterpreter()
	err := interp.Execute(`
		(set! a [1 2 3])
		(fmt.Println a)
		(fmt.Println a.0)
		(fmt.Println a.1)
		(fmt.Println a.2)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// [1 2 3]
	// 1
	// 2
	// 3
}

func ExampleInterpreter_Execute_map() {
	interp := lisp.NewInterpreter()
	err := interp.Execute(`
		(set! v {#a 1 #b 2 #c (+ 1 2)})
		(fmt.Println v)
		(fmt.Println v.0)
		(fmt.Println v.1)
		(fmt.Println v.2)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// {#a 1 #b 2 #c 3}
	// (#a 1)
	// (#b 2)
	// (#c 3)
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
	// (14 ())
}
