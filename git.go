package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func GetGitBranches(dir string) ([]string, error) {
	cmd := exec.Command("git", "branch", "--list")
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	branches := strings.Split(out.String(), "\n")
	var cleanBranches []string

	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if branch == "" {
			continue
		}
		prefix := strings.HasPrefix(branch, "* ")
		if prefix {
			branch = strings.TrimPrefix(branch, "* ")
		}
		cleanBranches = append(cleanBranches, branch)
	}
	return cleanBranches, nil
}

func GetGitLogs(dir string, branch string) (string, error) {
	cmd := exec.Command("git", "log", branch)
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return out.String(), err
}
