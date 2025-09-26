package compiler

type Program struct {
	Instructions []Instruction
	Counter      int
	State        map[string]int
	Labels       map[string]int // maps labels to instruction indices
	Snapshots    []Snapshot
}

func New(lines []string) Program {
	var program Program
	program.State = make(map[string]int)
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

func (p *Program) GetSnapshotAt(at int) (int, error) {
	if at < 0 || at >= len(p.Snapshots) {
		return -1, ErrProgramCounterOutOfBounds{Counter: at, Length: len(p.Snapshots)}
	}
	snapshot := p.Snapshots[at]
	return snapshot.Counter, nil
}

func (p *Program) SaveSnapshot() {
	snapshot := Snapshot{
		Counter: p.Counter,
		State:   make(map[string]int),
	}
	for k, v := range p.State {
		snapshot.State[k] = v
	}
	p.Snapshots = append(p.Snapshots, snapshot)
}
