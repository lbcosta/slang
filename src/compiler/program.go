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
}
