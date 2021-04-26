package ps

import (
	"errors"
	"fmt"

	"github.com/stack-machine-go/src/tk"
)

// Analyzes a list of instructions and returns the syntax errors it finds
func Parse(instructions *tk.Instructions) error {
	if instructions == nil {
		return errors.New("no instructions to analyze were provided")
	}

	for line, inst := range *instructions {
		switch true {
		case inst.Typ == tk.Push && inst.Val == nil:
			return fmt.Errorf("syntax error on line %d: PUSH instruction does not contain a valid number", line)
		case inst.Typ == tk.IfEq && inst.Jmp == nil:
			return fmt.Errorf("syntax error on line %d: IFEQ instruction does not contain a valid jump line number", line)
		case inst.Typ == tk.Jump && inst.Jmp == nil:
			return fmt.Errorf("syntax error on line %d: JUMP instruction does not contain a valid jump line number", line)
		}
	}
	return nil
}
