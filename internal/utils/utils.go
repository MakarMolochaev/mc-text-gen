package utils

import (
	"os"
	"strings"
)

func ReadAllLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	for i := range lines {
		lines[i] = strings.TrimSuffix(lines[i], "\r")
	}

	return lines, nil
}
