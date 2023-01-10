package common

import (
	"os"
	"strings"
)

func ReadInput() []string {
	if len(os.Args) != 2 {
		panic("wrong arguments")
	}
	return ReadInputFromFile(os.Args[1])
}

func ReadInputFromFile(name string) []string {
	bytes, err := os.ReadFile(name)
	Must(err)

	lines := strings.Split(string(bytes), "\n")

	return lines
}
