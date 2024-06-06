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
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	for err == nil {
		switch input {
		default: // unrecognized command
			fmt.Fprint(os.Stdout, strings.TrimSpace(input)+": command not found\n")
		}
		fmt.Fprint(os.Stdout, "$ ")
		input, err = bufio.NewReader(os.Stdin).ReadString('\n')
	}
}
