package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/stack-machine-go/src/in"
	"github.com/stack-machine-go/src/tk"
)

func main() {
	if len(os.Args) >= 2 {
		program, err := ioutil.ReadFile(os.Args[1])

		if err != nil {
			fmt.Println(err)
			return
		}

		instructions := tk.ParseExpression(string(program))

		interpreter, err := in.New(instructions)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = interpreter.Run()

		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("No file specified")
	}
}
