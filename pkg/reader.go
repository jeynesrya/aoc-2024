package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading file: %v", err)
	}

	return lines, nil
}
