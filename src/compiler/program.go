package compiler

const (
	OpAdd = iota
	OpSub
	OpIfNotEq
)

type Instruction struct {
	Label     string
	Operation int
	Args      []string
}

type Program struct {
	Instructions []Instruction
	Counter      int
	State        map[string]int
	LabelMap     map[string]int // maps labels to instruction indices
}

func New(lines []string) Program {
	var program Program
	program.State = make(map[string]int)
	program.Instructions = getInstructions(lines)
	program.newState()
	program.newLabelMap()
	program.Counter = 0

	return program
}

// newState reads all the instructions and initializes the state map with all variables set to 0.
// It assumes that the first argument of each instruction is always a variable.
func (p *Program) newState() {
	p.State = make(map[string]int)
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

// newLabelMap creates a map from labels to instruction indices for quick jumps.
func (p *Program) newLabelMap() {
	p.LabelMap = make(map[string]int)
	for i, instr := range p.Instructions {
		if instr.Label != "" {
			p.LabelMap[instr.Label] = i
		}
	}
}
