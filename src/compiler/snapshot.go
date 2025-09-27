package compiler

import (
	"fmt"
	"maps"
	"sort"
)

type Snapshot struct {
	counter     int
	state       map[string]int
	instruction Instruction
}

func (s Snapshot) String() string {
	return fmt.Sprintf("Counter: %d, State: %+v, Instruction: %s", s.counter, s.state, s.instruction.String())
}

func (s Snapshot) Print(withHeader bool) {
	if withHeader {
		s.PrintHeader()
	}

	// Collect all variable names in sorted order for consistent column order
	varNames := make([]string, 0, len(s.state))
	for varName := range s.state {
		varNames = append(varNames, varName)
	}

	// Sort variable names alphabetically
	sort.Strings(varNames)

	// Build state format string dynamically
	state := fmt.Sprintf("%-8d %-24s", s.counter, s.instruction.String())
	for _, varName := range varNames {
		state += fmt.Sprintf(" %-10d", s.state[varName])
	}

	fmt.Println(state)
}

func (s Snapshot) PrintHeader() {
	// Collect all variable names in sorted order for consistent column order
	varNames := make([]string, 0, len(s.state))
	for varName := range s.state {
		varNames = append(varNames, varName)
	}

	// Sort variable names alphabetically
	sort.Strings(varNames)

	// Build header format string dynamically
	header := fmt.Sprintf("%-8s %-24s", "Counter", "Instruction")
	for _, varName := range varNames {
		header += fmt.Sprintf(" %-10s", varName)
	}

	fmt.Println(header)
}

type Snapshots struct {
	snapshots []Snapshot
}

// PrintLast
func (s Snapshots) PrintLast(withHeader bool) {
	if len(s.snapshots) == 0 {
		fmt.Println("No snapshots available.")
		return
	}
	lastSnapshot := s.snapshots[len(s.snapshots)-1]
	lastSnapshot.Print(withHeader)
}

func (s Snapshots) GetSnapshotAt(at int) (int, error) {
	if at < 0 || at >= len(s.snapshots) {
		return -1, ErrProgramCounterOutOfBounds{Counter: at, Length: len(s.snapshots)}
	}
	snapshot := s.snapshots[at]
	return snapshot.counter, nil
}

func (s *Snapshots) SaveSnapshot(p *Program) {
	if s.snapshots == nil {
		s.snapshots = make([]Snapshot, 0)
	}

	// Create a deep copy of the program state to avoid mutation issues
	stateCopy := maps.Clone(p.State)

	snapshot := Snapshot{
		counter:     p.Counter,
		state:       stateCopy,
		instruction: p.Instructions[p.Counter],
	}

	s.snapshots = append(s.snapshots, snapshot)
}
