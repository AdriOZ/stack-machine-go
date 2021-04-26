package tk

// Represents an instruction of the program
type Token struct {
	// The type of token
	Typ Type

	// Value of the token. Only used for type Push.
	Val *float64

	// The jump to line value. Only used for type IfEq and Jump
	Jmp *uint32
}

// Shortcut to create a simple token
func NewToken(typ Type) Token {
	return Token{Typ: typ, Val: nil, Jmp: nil}
}

// Shortcut to create Push token
func NewPushToken(value float64) Token {
	return Token{Typ: Push, Val: &value, Jmp: nil}
}

// Shortcut to create an if token
func NewIfEqToken(jump uint32) Token {
	return Token{Typ: IfEq, Val: nil, Jmp: &jump}
}

// Shortcut to create a jump token
func NewJumpToken(jump uint32) Token {
	return Token{Typ: Jump, Val: nil, Jmp: &jump}
}
