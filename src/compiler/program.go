package compiler

import (
	"fmt"
	"log"
	"sort"
)

type Program struct {
	Instructions []Instruction
	Counter      int
	State        map[string]int
	Labels       map[string]int // maps labels to instruction indices
	Snapshots    []Snapshot
}

func New(lines []string) Program {
	var program Program
	program.Instructions = getInstructions(lines)
	program.newState()
	program.initLabels()
	program.Counter = 0

	return program
}

// newState reads all the instructions and initializes the state map with all variables set to 0.
// It assumes that the first argument of each instruction is always a variable.
func (p *Program) newState() {
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

func (p *Program) PrintStateHeader() {
	// Collect all variable names in sorted order for consistent column order
	varNames := make([]string, 0, len(p.State))
	for varName := range p.State {
		varNames = append(varNames, varName)
	}
	// Sort variable names alphabetically
	sort.Strings(varNames)

	// Build header format string dynamically
	header := fmt.Sprintf("%-8s %-20s", "Counter", "Instruction")
	for _, varName := range varNames {
		header += fmt.Sprintf(" %-10s", varName)
	}
	log.Println(header)
}

// PrintState prints the current state of the program in a formatted table row.
func (p *Program) PrintState() {
	// Collect all variable names in sorted order for consistent column order
	varNames := make([]string, 0, len(p.State))
	for varName := range p.State {
		varNames = append(varNames, varName)
	}
	// Sort variable names alphabetically
	sort.Strings(varNames)

	// Build state format string dynamically
	state := fmt.Sprintf("%-8d %-20s", p.Counter, p.Instructions[p.Counter].String())
	for _, varName := range varNames {
		state += fmt.Sprintf(" %-10d", p.State[varName])
	}
	log.Println(state)
}

func (p *Program) Run() error {
	// Pretty print a table header
	p.PrintStateHeader()
	p.PrintState()

	// 1. Saves a snapshot of the initial state.
	p.SaveSnapshot()

	// 2. Executes instructions until the counter goes out of bounds.
	for p.Counter < len(p.Instructions)-1 {
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
					// if label does not exist, halts program without error and explains why
					log.Printf("Label '%s' not found. Halting program.\n", label)
					return nil
				}
			} else {
				p.Counter++
			}
		default:
			return ErrInvalidInstruction{
				Details: fmt.Sprintf("unknown statement type: %v", instr.Statement),
				Line:    p.Counter + 1,
			}
		}

		// 3. Saves a snapshot of the current state.
		p.SaveSnapshot()
		// 4. Pretty prints the current state of the program.
		p.PrintState()
	}
	return nil
}
