package util

import (
	"fmt"
	"os"
	"path"
)

func LoadInput(day int, sample bool) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	inputPath := ""
	if sample {
		inputPath = path.Join(cwd, "sample", fmt.Sprintf("day%02d.txt", day))
	} else {
		inputPath = path.Join(cwd, "input", fmt.Sprintf("day%02d.txt", day))
	}
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("input file does not exist: " + inputPath)
	}
	out, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	return string(out)
}
