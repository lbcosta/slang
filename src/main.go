package main

import (
	"fmt"

	"lbcosta/slang/src/cmd"
	"lbcosta/slang/src/compiler"
)

func main() {
	programPath, programArgs, err := cmd.GetArgs()
	if err != nil {
		fmt.Printf("Error parsing arguments: %v\n", err)
		return
	}

	programLines, err := cmd.ReadLines(programPath)
	if err != nil {
		fmt.Printf("Error reading program file: %v\n", err)
		return
	}

	programInstructions := compiler.Compile(programLines)
	program := compiler.Build(programInstructions, programArgs)

	if err := program.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
	}

	fmt.Printf("\nResult: Y = %d\n", program.Output())
}
