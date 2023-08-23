package main

import (
	"bufio"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for {
		dir, err := os.Getwd()
		if err != nil {
			return
		}
		fmt.Print(dir, ":: ")
		if stdin.Scan() {
			cmdSlice := strings.Split(stdin.Text(), "|")
			actionController(cmdSlice)
		}
	}
}

func pwd() {
	workDir, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(workDir)
}

func cd(command []string) {
	if len(command) != 2 {
		fmt.Fprintln(os.Stderr, "Insert correct path")
	}
	err := os.Chdir(command[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func psFunc(...[]string) {
	fmt.Printf("%7s %-9s\n", "PID", "TTY")
	res, err := ps.Processes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for _, v := range res {
		fmt.Printf("%7d %-9s\n", v.Pid(), v.Executable())
	}
}

func echo(command []string) {
	if len(command) == 1 {
		fmt.Println("")
		return
	}
	for i := 1; i < len(command); i++ {
		command[i] = strings.Trim(command[i], "$")
		res := os.Getenv(command[i])
		fmt.Print(res, "\n")
	}
}

func kill(command []string) {
	if len(command) != 2 {
		fmt.Fprintln(os.Stderr, "Insert correct command")
	}
	pid, err := strconv.Atoi(command[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	proc.Kill()
}

func actionController(data []string) {
	for _, command := range data {
		commandSl := strings.Split(command, " ")
		switch commandSl[0] {
		case "pwd":
			pwd()
		case "ps":
			psFunc(commandSl)
		case "kill":
			kill(commandSl)
		case "cd":
			cd(commandSl)
		case "echo":
			echo(commandSl)
		}
	}
}
