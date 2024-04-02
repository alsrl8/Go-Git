package core

import (
	"Go-Git/ai"
	"Go-Git/config"
	"Go-Git/gitlog"
	"Go-Git/terminal"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Service(conf config.Config) {
	terminal.ClearTerminal()

	rootDir := conf.GetRootDirectory()
	branch := selectBranch(rootDir)
	for {
		gitLogNum := inputGitLogNumber()
		gitLogs, _ := gitlog.GetGitLogs(rootDir, branch, gitLogNum)

		commits := gitlog.ParseGitLogs(gitLogs)
		confirm := confirmSummarize(commits)
		if !confirm {
			continue
		}
		ai.RequestToSummarizeGitLogs(commits)
		break
	}
}

func selectBranch(rootDir string) (userBranch string) {
	branches := gitlog.GetGitBranches(rootDir)

	reader := bufio.NewReader(os.Stdin)

	alert := ""
	valid := false
	terminal.ClearTerminal()
	for !valid {
		if alert != "" {
			terminal.PrintAlert(alert)
		}
		s := fmt.Sprintf("branches: %v\n", branches)
		terminal.PrintNotice(s)
		terminal.PrintNotice("Enter a branch from the list: ")
		userBranch, _ = reader.ReadString('\n')
		userBranch = strings.TrimSpace(userBranch)

		for _, branch := range branches {
			if branch == userBranch {
				valid = true
				break
			}
		}

		if !valid {
			alert = "Invalid branch. Please enter a branch from the list.\n"
		}
	}
	return userBranch
}

func inputGitLogNumber() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		terminal.PrintNotice("Enter number of git logs: ")
		userInput, _ := reader.ReadString('\n')
		num, err := strconv.Atoi(strings.Trim(userInput, "\r\n"))
		if err != nil {
			terminal.PrintAlert("Wrong number\n")
			continue
		}
		return num
	}
}

func confirmSummarize(commits []gitlog.Commit) bool {
	for _, commit := range commits {
		fmt.Printf("%v\n", commit)
	}
	alert := ""
	for {
		if alert != "" {
			terminal.PrintAlert(alert)
		}
		terminal.PrintNotice("Are you gonna summarize this git commits? : (Y/N)")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\r\n")

		switch input {
		case "Y", "y":
			return true
		case "N", "n":
			return false
		default:
			alert = "Wrong input. Choose one of them (Y/M)\n"
			continue
		}
	}

}
