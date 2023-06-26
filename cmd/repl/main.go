package main

import (
	"fmt"
	"log"

	"github.com/aziis98/go-lisp-interpreter"
	"github.com/chzyer/readline"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	interp := lisp.NewInterpreter()

	rl, err := readline.New("=> ")
	if err != nil {
		log.Fatal(err)
	}

	for {
		line, err := rl.Readline()
		if err != nil {
			log.Fatal(err)
		}

		v, err := interp.Eval(line)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(lisp.Show(v))
		}
	}
}
