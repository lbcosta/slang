package compiler

import "fmt"

// Possible statements
const (
	Increment = iota
	Decrement
	ConditionalBranch
	Halt
)

type Instruction struct {
	Label     string
	Statement int
	Args      []string
}

func (instr Instruction) String() string {
	prefix := "    "
	if instr.Label != "" {
		prefix = fmt.Sprintf("[%s] ", instr.Label)
	}

	switch instr.Statement {
	case Increment:
		return fmt.Sprintf("%s%s <- %s + 1", prefix, instr.Args[0], instr.Args[0])
	case Decrement:
		return fmt.Sprintf("%s%s <- %s - 1", prefix, instr.Args[0], instr.Args[0])
	case ConditionalBranch:
		return fmt.Sprintf("%sIF %s != 0 GOTO %s", prefix, instr.Args[0], instr.Args[1])
	case Halt:
		return fmt.Sprintf("%sHALT", prefix)
	default:
		return "UNKNOWN INSTRUCTION"
	}
}
