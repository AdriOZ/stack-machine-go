package tk

// Type of a token
type Type = uint8

const (
	Push      Type = 0
	Pop       Type = 1
	Add       Type = 2
	IfEq      Type = 3
	Jump      Type = 4
	PrintNumb Type = 5
	PrintChar Type = 6
	Duplicate Type = 7
	Mul       Type = 8
	Div       Type = 9
	Swap      Type = 10
	Sub       Type = 11
)
