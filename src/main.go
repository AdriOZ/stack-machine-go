package main

import (
	"fmt"
	"io/ioutil"

	"github.com/stack-machine-go/src/in"
	"github.com/stack-machine-go/src/tk"
)

func main() {
	program, _ := ioutil.ReadFile("test.stack")
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

}
