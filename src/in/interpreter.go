package in

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
	"github.com/stack-machine-go/src/ps"
	"github.com/stack-machine-go/src/tk"
)

// Interpreter of the stack machine
type Interpreter struct {
	instructions *tk.Instructions
	index        uint32
	lenght       uint32
	stack        *stack.Stack
}

// Creates a new interpreter for the list of expressions
func New(instructions *tk.Instructions) (*Interpreter, error) {
	err := ps.Parse(instructions)

	if err != nil {
		return nil, err
	}

	min := -1
	max := 0

	for index := range *instructions {
		if min == -1 || int(index) < min {
			min = int(index)
		}
		if int(index) > max {
			max = int(index)
		}
	}

	return &Interpreter{
		instructions: instructions,
		index:        uint32(min),
		lenght:       uint32(max) + 1,
		stack:        stack.New(),
	}, nil
}

// Interprets the instructions of a program.
// Returns an error if a problem happens during the execution
func (inter *Interpreter) Run() error {
	for inter.index < inter.lenght {
		current, exists := (*inter.instructions)[inter.index]
		advance := true

		if exists {
			switch current.Typ {
			case tk.Push:
				inter.stack.Push(*current.Val)
			case tk.Pop:
				inter.stack.Pop()
			case tk.Add:
				if inter.stack.Len() >= 2 {
					num1 := inter.stack.Pop().(float64)
					num2 := inter.stack.Pop().(float64)
					inter.stack.Push(num1 + num2)
				}
			case tk.Mul:
				if inter.stack.Len() >= 2 {
					num1 := inter.stack.Pop().(float64)
					num2 := inter.stack.Pop().(float64)
					inter.stack.Push(num1 * num2)
				}
			case tk.Div:
				if inter.stack.Len() >= 2 {
					num1 := inter.stack.Pop().(float64)
					num2 := inter.stack.Pop().(float64)
					inter.stack.Push(num1 / num2)
				}
			case tk.Swap:
				if inter.stack.Len() >= 2 {
					num1 := inter.stack.Pop().(float64)
					num2 := inter.stack.Pop().(float64)
					inter.stack.Push(num1)
					inter.stack.Push(num2)
				}
			case tk.IfEq:
				if inter.stack.Len() > 0 && inter.stack.Peek().(float64) != 0.0 {
					if *current.Jmp < 1 || *current.Jmp > inter.lenght {
						return fmt.Errorf("\nerror on line %d IFEQ trying to jump to the non-existent line %d", inter.index, *current.Jmp)
					}
					inter.index = *current.Jmp
					advance = false
				}
			case tk.Jump:
				if *current.Jmp < 1 || *current.Jmp > inter.lenght {
					return fmt.Errorf("\nerror on line %d JUMP trying to jump to the non-existent line %d", inter.index, *current.Jmp)
				}
				inter.index = *current.Jmp
				advance = false
			case tk.PrintNumb:
				if inter.stack.Len() > 0 {
					fmt.Printf("%f", inter.stack.Peek().(float64))
				}
			case tk.PrintChar:
				if inter.stack.Len() > 0 {
					fmt.Print(string(rune(int(inter.stack.Peek().(float64)))))
				}
			case tk.Duplicate:
				if inter.stack.Len() > 0 {
					inter.stack.Push(inter.stack.Peek())
				}
			}
		}
		if advance {
			inter.index += 1
		}
	}
	return nil
}
