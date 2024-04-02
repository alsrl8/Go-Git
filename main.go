package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	configFilename := "./gogit.config"
	config, err := LoadConfig(configFilename)
	if err != nil {
		fmt.Printf("Error loading config: %s\n", err)
		return
	}

	config.PrintConfig()
	rootDir, err := config.GetRootDirectory()
	if err != nil {
		fmt.Printf("Error getting root directory: %s\n", err)
		return
	}

	branches, err := GetGitBranches(rootDir)
	if err != nil {
		fmt.Printf("Error getting branches: %s\n", err)
		return
	}

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
			ClearTerminal()
			fmt.Println("Invalid branch. Please enter a branch from the list.")
			fmt.Printf("branches :%v\n", branches)
		}
	}

	fmt.Printf("Selected branch: %s\n", userBranch)

	logs, err := GetGitLogs(rootDir, userBranch)
	if err != nil {
		fmt.Printf("Error getting logs from directory: %s\n", err)
		return
	}

	fmt.Printf("logs: \n%s\n", logs)
}
