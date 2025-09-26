package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"lbcosta/slang/src/compiler"
	"lbcosta/slang/src/util"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: slang PROGRAM [VAR=VALUE ...]")
		return
	}

	lines, err := util.ReadLines(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var initialVars map[string]int
	if len(os.Args) > 2 {
		initialVars = make(map[string]int)
		for _, arg := range os.Args[2:] {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) != 2 {
				fmt.Printf("Invalid argument format: %s (expected VAR=VALUE)\n", arg)
				return
			}

			name := strings.TrimSpace(parts[0])
			if name == "" {
				fmt.Printf("Invalid variable name in argument: %s\n", arg)
				return
			}
			name = strings.ToUpper(name)

			valueStr := strings.TrimSpace(parts[1])
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				fmt.Printf("Invalid value for %s: %v\n", name, err)
				return
			}
			if value < 0 {
				fmt.Printf("Invalid value for %s: must be non-negative\n", name)
				return
			}
			initialVars[name] = value
		}
	}

	program := compiler.New(lines, initialVars)

	if err := program.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
	}

	program.PrintResult()
}
