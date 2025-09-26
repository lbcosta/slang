package compiler

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
