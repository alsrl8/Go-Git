package gitlog

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// IsGitRepository checks if a given path is a git repository.
func IsGitRepository(path string) bool {
	cmd := exec.Command("git", "branch", "--list")
	cmd.Dir = path

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

// GetGitBranches retrieves the list of branches in a given git repository directory.
func GetGitBranches(dir string) []string {
	cmd := exec.Command("git", "branch", "--list")
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	_ = cmd.Run()

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
	return cleanBranches
}

// GetGitLogs retrieves the git logs for a given directory, branch, and number of logs.
func GetGitLogs(dir string, branch string, num int) (string, error) {
	fmtStr := "%H%n%an%n%ae%n%ad%n%B%n%%%"
	fmtOption := fmt.Sprintf("--pretty=format:%s", fmtStr)

	cmd := exec.Command("git", "log", "-n", strconv.Itoa(num), branch, fmtOption)
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return out.String(), err
}
