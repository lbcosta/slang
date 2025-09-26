package compiler

import (
	"fmt"
)

type ErrInvalidInstruction struct {
	Line    int
	Details string
}

func (e ErrInvalidInstruction) Error() string {
	return fmt.Sprintf("invalid instruction at line %d: %s", e.Line, e.Details)
}

type ErrProgramCounterOutOfBounds struct {
	Counter int
	Length  int
}

func (e ErrProgramCounterOutOfBounds) Error() string {
	return fmt.Sprintf("program counter %d out of bounds (0 to %d)", e.Counter, e.Length-1)
}
