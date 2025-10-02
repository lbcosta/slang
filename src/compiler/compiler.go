package compiler

import (
	"regexp"
	"strings"
)

// Valid instructions:
// V <- V + 1
// V <- V - 1
// IF V != 0 GOTO L

// V is a variable (string). Its value is always a non-negative integer.
// L is a label (string). It maps to an instruction index in the program.
// Every instruction can optionally start with a label [L].
// V and L are both exactly one uppercase letter with optional digits, e.g., X1, A, B1, C2, etc.

// Examples:
// [A] V <- V + 1
// [B1] V <- V - 1
// [C1] IF V != 0 GOTO L

// Every text is always in uppercase and trimmed of spaces.
// Lines that are empty or start with # are ignored.

// lineToInstruction converts a line of text to an Instruction struct (Label, Operation, Args).
func lineToInstruction(line string) (Instruction, error) {
	var instr Instruction
	parts := strings.Fields(line)

	// Check for label
	if strings.HasPrefix(parts[0], "[") && strings.HasSuffix(parts[0], "]") {
		instr.Label = strings.Trim(parts[0], "[]")
		parts = parts[1:]
	}

	// Parse operation with regex of three possible formats
	// V <- V + 1
	reIncr := regexp.MustCompile(`^([A-Z][0-9]*)\s*<-\s*([A-Z][0-9]*)\s*\+\s*1$`)
	// V <- V - 1
	reDecr := regexp.MustCompile(`^([A-Z][0-9]*)\s*<-\s*([A-Z][0-9]*)\s*-\s*1$`)
	// IF V != 0 GOTO L
	reIfNotEq := regexp.MustCompile(`^IF\s+([A-Z][0-9]*)\s*!=\s*0\s*GOTO\s+([A-Z][0-9]*)$`)
	// W <- f(V1, V2, ..., Vn)  (macro for assignment)
	reMacroAssign := regexp.MustCompile(`^([A-Z][0-9]*)\s*<-\s*([A-Z]+)\((([A-Z][0-9]*(,\s*)?)*)\)$`)
	// IF f(V1, V2, ..., Vn) != 0 GOTO L  (macro for conditional branch)
	reMacroCond := regexp.MustCompile(`^IF\s+([A-Z]+)\((([A-Z][0-9]*(,\s*)?)*)\)\s*!=\s*0\s*GOTO\s+([A-Z][0-9]*)$`)

	lineStr := strings.Join(parts, " ")

	if matches := reIncr.FindStringSubmatch(lineStr); matches != nil {
		// matches[1] and matches[2] should be the same variable
		if matches[1] == matches[2] {
			instr.Statement = Increment
			instr.Args = []string{matches[1]} // V
			return instr, nil
		}
	} else if matches := reDecr.FindStringSubmatch(lineStr); matches != nil {
		// matches[1] and matches[2] should be the same variable
		if matches[1] == matches[2] {
			instr.Statement = Decrement
			instr.Args = []string{matches[1]} // V
			return instr, nil
		}
	} else if matches := reIfNotEq.FindStringSubmatch(lineStr); matches != nil {
		instr.Statement = ConditionalBranch
		instr.Args = []string{matches[1], matches[2]} // V, L
		return instr, nil
	} else if matches := reMacroAssign.FindStringSubmatch(lineStr); matches != nil {
		// matches[1]: W, matches[2]: f, matches[3]: V1, V2, ..., Vn (comma separated)
		instr.Statement = AssignmentMacro
		args := []string{matches[1], matches[2]}
		// Split the argument list by comma and trim spaces
		if matches[3] != "" {
			for _, v := range strings.Split(matches[3], ",") {
				v = strings.TrimSpace(v)
				if v != "" {
					args = append(args, v)
				}
			}
		}
		instr.Args = args
		return instr, nil
	} else if matches := reMacroCond.FindStringSubmatch(lineStr); matches != nil {
		// matches[1]: f, matches[2]: V1, V2, ..., Vn (comma separated), matches[5]: L
		instr.Statement = ConditionalMacro
		args := []string{matches[1]}
		// Split the argument list by comma and trim spaces
		if matches[2] != "" {
			for _, v := range strings.Split(matches[2], ",") {
				v = strings.TrimSpace(v)
				if v != "" {
					args = append(args, v)
				}
			}
		}
		args = append(args, matches[5]) // L
		instr.Args = args
		return instr, nil
	}

	return Instruction{}, ErrInvalidInstruction{}
}

func Compile(programLines []string) []Instruction {
	var instructions []Instruction
	for lineIdx, line := range programLines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue // Skip empty lines and comments
		}
		instr, err := lineToInstruction(line)
		if err != nil {
			panic(ErrInvalidInstruction{Line: lineIdx + 1, Details: line})
		}

		instructions = append(instructions, instr)
	}

	haltInstr := Instruction{Statement: Halt}
	instructions = append(instructions, haltInstr)

	return instructions
}
