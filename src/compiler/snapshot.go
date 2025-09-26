package compiler

import "fmt"

type Snapshot struct {
	Counter int
	State   map[string]int
}

func (s Snapshot) String() string {
	return fmt.Sprintf("Counter: %d, State: %+v", s.Counter, s.State)
}
