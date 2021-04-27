package tk

import (
	"strconv"
	"strings"
)

// List of instructions
// where the index of the map is the line where the expression appears
// and the value is the instruction
type Instructions map[uint32]Token

// Given a the source code of the program, returns a list of instructions
// where the index of the map is the line where the expression appears
// and the value is the instruction
func ParseExpression(program string) *Instructions {
	lines := strings.Split(strings.ToUpper(program), "\n")
	result := make(Instructions)

	for index, line := range lines {
		trimmed := strings.TrimSpace(line)
		switch true {
		case strings.HasPrefix(trimmed, "PUSH"):
			parts := strings.Split(trimmed, " ")

			if len(parts) == 1 {
				result[uint32(index+1)] = NewToken(Push)
			} else {
				val, err := strconv.ParseFloat(parts[1], 64)

				if err != nil {
					result[uint32(index+1)] = NewToken(Push)
				} else {
					result[uint32(index+1)] = NewPushToken(val)
				}
			}
		case strings.HasPrefix(trimmed, "POP"):
			result[uint32(index+1)] = NewToken(Pop)
		case strings.HasPrefix(trimmed, "ADD"):
			result[uint32(index+1)] = NewToken(Add)
		case strings.HasPrefix(trimmed, "MUL"):
			result[uint32(index+1)] = NewToken(Mul)
		case strings.HasPrefix(trimmed, "DIV"):
			result[uint32(index+1)] = NewToken(Div)
		case strings.HasPrefix(trimmed, "SWAP"):
			result[uint32(index+1)] = NewToken(Swap)
		case strings.HasPrefix(trimmed, "IFEQ"):
			parts := strings.Split(trimmed, " ")

			if len(parts) == 1 {
				result[uint32(index+1)] = NewToken(IfEq)
			} else {
				val, err := strconv.ParseUint(parts[1], 10, 32)

				if err != nil {
					result[uint32(index+1)] = NewToken(IfEq)
				} else {
					result[uint32(index+1)] = NewIfEqToken(uint32(val))
				}
			}
		case strings.HasPrefix(trimmed, "JUMP"):
			parts := strings.Split(trimmed, " ")

			if len(parts) == 1 {
				result[uint32(index+1)] = NewToken(Jump)
			} else {
				val, err := strconv.ParseUint(parts[1], 10, 32)

				if err != nil {
					result[uint32(index+1)] = NewToken(Jump)
				} else {
					result[uint32(index+1)] = NewJumpToken(uint32(val))
				}
			}
		case strings.HasPrefix(trimmed, "PRINTN"):
			result[uint32(index+1)] = NewToken(PrintNumb)
		case strings.HasPrefix(trimmed, "PRINTC"):
			result[uint32(index+1)] = NewToken(PrintChar)
		case strings.HasPrefix(trimmed, "DUP"):
			result[uint32(index+1)] = NewToken(Duplicate)
		default:
		}
	}
	return &result
}
