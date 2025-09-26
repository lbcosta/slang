package compiler

import "fmt"

type Snapshot struct {
	Counter int
	State   map[string]int
}

func (s Snapshot) String() string {
	return fmt.Sprintf("Counter: %d, State: %+v", s.Counter, s.State)
}

// Program:

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
