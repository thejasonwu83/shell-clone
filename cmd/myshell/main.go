package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	switch input {
	default: // unrecognized command
		fmt.Fprint(os.Stdout, strings.TrimSpace(input)+": command not found\n")
	}
}
