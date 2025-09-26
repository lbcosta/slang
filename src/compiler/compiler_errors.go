package compiler

import (
	"errors"
)

// Invalid instruction error. Shows the line number and content.
var ErrInvalidInstruction = errors.New("invalid instruction")
