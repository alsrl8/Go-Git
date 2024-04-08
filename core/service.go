package core

import (
	"Go-Git/ai"
	"Go-Git/config"
	"Go-Git/gitlog"
	"Go-Git/slack"
	"Go-Git/terminal"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Service(conf config.Config) {
	terminal.ClearTerminal()

	rootDir := conf.GetRootDirectory()
	branch := selectBranch(rootDir)
	var summary string
	for {
		gitLogNum := inputGitLogNumber()
		gitLogs, _ := gitlog.GetGitLogs(rootDir, branch, gitLogNum)

		commits := gitlog.ParseGitLogs(gitLogs)
		reports, confirm := confirmSummarize(commits)
		if !confirm {
			continue
		}
		summary = ai.RequestToSummarizeGitLogs(commits, reports)
		break
	}

	channelId := "D06E3025AF6"
	messages, err := slack.FetchChannelMessages(channelId)
	if err != nil {
		log.Fatalf("Failed to fetch channel messages from slack api. %v\n", err)
		return
	}

	// TODO Channel Message 원하는 걸 못 찾았다면 어떻게 처리할지?

	ts := messages.Messages[0].Ts
	slack.PostMessage(channelId, summary, ts)
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

func confirmSummarize(commits []gitlog.Commit) ([]string, bool) {
	reader := bufio.NewReader(os.Stdin)

	for _, commit := range commits {
		fmt.Printf("%v\n", commit)
	}

	terminal.PrintNotice("If you have something more to report, enter it. After entering all reports, please enter `q`\n")
	var reports []string
InputReport:
	for {
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\r\n")

		switch input {
		case "q", "Q":
			break InputReport
		default:
			reports = append(reports, input)
		}
	}

	alert := ""
	for {
		if alert != "" {
			terminal.PrintAlert(alert)
		}
		terminal.PrintNotice("Are you gonna summarize this git commits? : (Y/N)")
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\r\n")

		switch input {
		case "Y", "y":
			return reports, true
		case "N", "n":
			return reports, false
		default:
			alert = "Wrong input. Choose one of them (Y/M)\n"
			continue
		}
	}

}
