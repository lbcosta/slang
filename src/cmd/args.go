package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetArgs() (programPath string, programArgs map[string]int, err error) {
	if len(os.Args) < 2 {
		fmt.Println("Usage: slang PROGRAM [VAR=VALUE ...]")
		return "", nil, fmt.Errorf("no program path provided")
	}

	programPath = os.Args[1]
	if len(programPath) == 0 {
		return "", nil, fmt.Errorf("no program path provided")
	}

	// Read "programPath" and add suffix ".slang" if not present
	if len(programPath) < 5 || programPath[len(programPath)-5:] != ".slang" {
		programPath += ".slang"
	}

	if len(os.Args) > 2 {
		programArgs = make(map[string]int)
		for _, arg := range os.Args[2:] {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) != 2 {
				return "", nil, fmt.Errorf("invalid argument format: %s, expected VAR=VALUE", arg)
			}

			name := strings.TrimSpace(parts[0])
			if name == "" {
				return "", nil, fmt.Errorf("invalid variable name in argument: %s", arg)
			}
			name = strings.ToUpper(name)

			valueStr := strings.TrimSpace(parts[1])
			value, atoiErr := strconv.Atoi(valueStr)
			if atoiErr != nil {
				return "", nil, fmt.Errorf("invalid value for %s: %v", name, atoiErr)
			}
			if value < 0 {
				return "", nil, fmt.Errorf("invalid value for %s: must be non-negative", name)
			}
			programArgs[name] = value
		}
	}

	return programPath, programArgs, nil
}

// ReadLines reads a file and returns a slice of trimmed lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
