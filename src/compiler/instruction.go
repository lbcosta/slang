package compiler

import "fmt"

// Possible statements
const (
	Increment = iota
	Decrement
	ConditionalBranch
	Halt
	AssignmentMacro
	ConditionalMacro
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
	case AssignmentMacro:
		// args[0]: W, args[1]: f, args[2...]: V1, V2, ..., Vn
		W := instr.Args[0]
		f := instr.Args[1]
		Vs := instr.Args[2:]
		return fmt.Sprintf("%s%s <- %s(%s)", prefix, W, f, joinArgs(Vs))
	case ConditionalMacro:
		// args[0]: f, args[1...n-1]: V1, V2, ..., Vn, args[n]: L
		f := instr.Args[0]
		Vs := instr.Args[1 : len(instr.Args)-1]
		L := instr.Args[len(instr.Args)-1]
		return fmt.Sprintf("%sIF %s(%s) != 0 GOTO %s", prefix, f, joinArgs(Vs), L)
	default:
		return "UNKNOWN INSTRUCTION"
	}
}

func joinArgs(args []string) string {
	result := ""
	for i, arg := range args {
		if i > 0 {
			result += ", "
		}
		result += arg
	}
	return result
}
