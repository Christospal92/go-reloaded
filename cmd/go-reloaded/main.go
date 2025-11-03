package main

import (
	"fmt"
	"os"
	"strings"

	"go-reloaded/pkg/transform"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go-reloaded <input.txt> <output.txt>")
		os.Exit(1)
	}

	inPath := os.Args[1]
	outPath := os.Args[2]

	data, err := os.ReadFile(inPath)
	if err != nil {
		fmt.Println("error reading input:", err)
		os.Exit(1)
	}

	out := transform.Transform(string(data))

	out = strings.TrimRight(out, "\r\n") + "\n"

	if err := os.WriteFile(outPath, []byte(out), 0o644); err != nil {
		fmt.Println("error writing output:", err)
		os.Exit(1)
	}
}
