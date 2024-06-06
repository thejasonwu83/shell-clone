package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getArgs(input string) []string {
	input = strings.TrimSpace(input)
	args := []string{}
	for input != "" {
		argEnd := strings.Index(input, " ")
		if argEnd == -1 { // reached last argument
			args = append(args, input)
			break
		}
		args = append(args, input[:argEnd])
		input = strings.TrimSpace(input[argEnd:])
	}
	return args
}

func echo(args []string) {
	msg := ""
	for _, arg := range args {
		msg += arg + " "
	}
	msg = strings.TrimSpace(msg)
	fmt.Fprint(os.Stdout, msg+"\n")
}

func getType(args []string) {
	for _, arg := range args {
		arg = strings.ToLower(arg)
		msg := ""
		switch arg {
		case "echo":
			msg = "echo is a shell builtin\n"
		case "exit":
			msg = "exit is a shell builtin\n"
		case "type":
			msg = "type is a shell builtin\n"
		default:
			msg = arg + " not found\n"
		}
		fmt.Fprint(os.Stdout, msg)
	}
}

// TODO: check for valid num args

func evalCmd(input string) {
	args := getArgs(input)
	switch strings.ToLower(args[0]) {
	case "exit":
		if len(args) > 1 {
			exitStatus, _ := strconv.Atoi(args[1])
			os.Exit(exitStatus)
		}
		os.Exit(0)
	case "echo":
		echo(args[1:])
	case "type":
		getType(args[1:])
	default: // unrecognized command
		fmt.Fprint(os.Stdout, strings.TrimSpace(input)+": command not found\n")
	}
}

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// REPL loop
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	for err == nil {
		evalCmd(input)
		fmt.Fprint(os.Stdout, "$ ")
		input, err = bufio.NewReader(os.Stdin).ReadString('\n')
	}
}
