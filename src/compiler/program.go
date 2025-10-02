package compiler

import (
	"fmt"
)

type Program struct {
	Instructions []Instruction
	Counter      int
	State        map[string]int
	Labels       map[string]int // maps labels to instruction indices
	Snapshots    Snapshots
	Macros       map[string]Program // maps macro names to their programs
}

func Build(instructions []Instruction, initialState map[string]int) Program {
	var program Program
	program.Instructions = instructions
	program.newState(initialState)
	program.initLabels()
	program.Counter = 0

	return program
}

// newState reads all the instructions and initializes the state map with all variables set to 0.
// It assumes that the first argument of each instruction is always a variable.
func (p *Program) newState(initialState map[string]int) {
	p.State = make(map[string]int)
	p.State["Y"] = 0 // Ensure "Y" is always initialized

	for _, instr := range p.Instructions {
		// Initialize variables in state map. The first arg is always a variable.
		if len(instr.Args) > 0 {
			varName := instr.Args[0]
			if _, exists := p.State[varName]; !exists {
				p.State[varName] = 0
			}
		}
		// For OpIfNotEq, the second arg is a label, so we don't initialize it.
	}

	for varName, value := range initialState {
		p.State[varName] = value
	}
}

// initLabels creates a map from labels to instruction indices for quick jumps.
func (p *Program) initLabels() {
	p.Labels = make(map[string]int)
	for i, instr := range p.Instructions {
		if instr.Label != "" {
			p.Labels[instr.Label] = i
		}
	}
}

func (p *Program) Length() int {
	return len(p.Instructions)
}

func (p *Program) Run() error {
	p.Snapshots.SaveSnapshot(p)
	p.Snapshots.PrintLast(true)

	// 2. Executes instructions until the counter goes out of bounds.
	for p.Counter < len(p.Instructions) {
		instr := p.Instructions[p.Counter]
		switch instr.Statement {
		case Increment:
			varName := instr.Args[0]
			p.State[varName]++
			p.Counter++
		case Decrement:
			varName := instr.Args[0]
			if p.State[varName] > 0 {
				p.State[varName]--
			}
			p.Counter++
		case ConditionalBranch:
			varName := instr.Args[0]
			label := instr.Args[1]
			if p.State[varName] != 0 {
				if target, exists := p.Labels[label]; exists {
					p.Counter = target
				} else {
					// if label does not exist, program should halt
					p.Counter = p.Length() - 1
				}
			} else {
				p.Counter++
			}
		case Halt:
			// Halts the program gracefully
			return nil
		case AssignmentMacro:
			// args[0]: W, args[1]: f, args[2...]: V1, V2, ..., Vn
			if len(instr.Args) < 3 {
				return ErrInvalidInstruction{
					Details: "not enough arguments for AssignmentMacro",
					Line:    p.Counter + 1,
				}
			}

			fmt.Println("Executing AssignmentMacro:", instr)
			return nil // Placeholder for actual implementation
		case ConditionalMacro:
			// args[0]: f, args[1...n-1]: V1, V2, ..., Vn, args[n]: L
			if len(instr.Args) < 3 {
				return ErrInvalidInstruction{
					Details: "not enough arguments for ConditionalMacro",
					Line:    p.Counter + 1,
				}
			}
			f := instr.Args[0]
			Vs := instr.Args[1 : len(instr.Args)-1]
			L := instr.Args[len(instr.Args)-1]

			println("Executing ConditionalMacro:", f, Vs, L)
			return nil // Placeholder for actual implementation
		default:
			return ErrInvalidInstruction{
				Details: fmt.Sprintf("unknown statement type: %v", instr.Statement),
				Line:    p.Counter + 1,
			}
		}

		p.Snapshots.SaveSnapshot(p)
		p.Snapshots.PrintLast(false)
	}

	return nil
}

func (s *Program) Output() int {
	return s.State["Y"]
}
