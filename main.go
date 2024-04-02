package main

import (
	"Go-Git/config"
	"Go-Git/gitlog"
	"Go-Git/terminal"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	configFilename := "./gogit.config"
	config_ := config.SetConfig(configFilename)

	rootDir := config_.GetRootDirectory()
	branches := gitlog.GetGitBranches(rootDir)

	fmt.Printf("branches :%v\n", branches)

	reader := bufio.NewReader(os.Stdin)

	valid := false
	var userBranch string
	for !valid {
		fmt.Print("Enter a branch from the list: ")
		userBranch, _ = reader.ReadString('\n')
		userBranch = strings.TrimSpace(userBranch)

		for _, branch := range branches {
			if branch == userBranch {
				valid = true
				break
			}
		}

		if !valid {
			terminal.ClearTerminal()
			fmt.Println("Invalid branch. Please enter a branch from the list.")
			fmt.Printf("branches :%v\n", branches)
		}
	}

	fmt.Printf("Selected branch: %s\n", userBranch)

	valid = false
	var gitLogNum int
	for !valid {
		fmt.Print("Enter number of git logs: ")
		userInput, _ := reader.ReadString('\n')
		num, err := strconv.Atoi(strings.Trim(userInput, "\r\n"))
		if err != nil {
			fmt.Println("Wrong number")
			continue
		}
		gitLogNum = num
		valid = true
	}

	logs, err := gitlog.GetGitLogs(rootDir, userBranch, gitLogNum)
	if err != nil {
		fmt.Printf("Error getting logs from directory: %s\n", err)
		return
	}

	fmt.Printf("logs: \n%s\n", logs)
}
