package main

import (
	"fmt"
	"os"

	"lbcosta/slang/src/compiler"
	"lbcosta/slang/src/util"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: slang PROGRAM")
		return
	}

	lines, err := util.ReadLines(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	program := compiler.ParseProgram(lines)

	// print program instructions
	for i, instr := range program.Instructions {
		fmt.Printf("%d: %+v\n", i, instr)
	}

	// print counter
	fmt.Printf("Counter: %d\n", program.Counter)

	// print initial state
	fmt.Printf("Initial State: %+v\n", program.State)

	// Here you would typically run the program or further process it
}
