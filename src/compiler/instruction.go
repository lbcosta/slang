package compiler

import "fmt"

// Possible statements
const (
	Increment = iota
	Decrement
	ConditionalBranch
)

type Instruction struct {
	Label     string
	Statement int
	Args      []string
}

func (instr Instruction) String() string {
	switch instr.Statement {
	case Increment:
		return fmt.Sprintf("%s <- %s + 1", instr.Args[0], instr.Args[0])
	case Decrement:
		return fmt.Sprintf("%s <- %s - 1", instr.Args[0], instr.Args[0])
	case ConditionalBranch:
		return fmt.Sprintf("IF %s != 0 GOTO %s", instr.Args[0], instr.Args[1])
	default:
		return "UNKNOWN INSTRUCTION"
	}
}
