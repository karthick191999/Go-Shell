package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	//fmt.Println("Just another print stmt")
	reader := bufio.NewReader(os.Stdin)
	for {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		currentUser, err := user.Current()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Printf("%s@%s-%s>", currentUser.Username, hostname, currentDir)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := executeCommands(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
func executeCommands(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			// need to go the home directory
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			return os.Chdir(homeDir)
		}
		return os.Chdir(args[1])
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()

}
