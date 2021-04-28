package tk

// Type of a token
type Type = uint8

const (
	Push      Type = 0
	Pop       Type = 1
	Duplicate Type = 2
	Swap      Type = 3
	Add       Type = 4
	Sub       Type = 5
	Mul       Type = 6
	Div       Type = 7
	IfEq      Type = 8
	Jump      Type = 9
	PrintNumb Type = 10
	PrintChar Type = 11
	PrintInt  Type = 12
)
