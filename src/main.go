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

	program := compiler.New(lines)

	if err := program.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
	}
}
