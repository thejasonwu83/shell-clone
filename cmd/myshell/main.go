package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	for err == nil {
		args := strings.Split(strings.TrimSpace(input), " ")
		switch strings.ToLower(args[0]) {
		case "exit":
			if len(args) > 1 {
				exitStatus, _ := strconv.Atoi(args[1])
				os.Exit(exitStatus)
			}
			os.Exit(0)
		default: // unrecognized command
			fmt.Fprint(os.Stdout, strings.TrimSpace(input)+": command not found\n")
		}
		fmt.Fprint(os.Stdout, "$ ")
		input, err = bufio.NewReader(os.Stdin).ReadString('\n')
	}
}
