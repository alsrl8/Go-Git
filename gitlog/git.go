package gitlog

import (
	"bytes"
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

func GetGitLogs(dir string, branch string, num int) (string, error) {
	cmd := exec.Command("git", "log", "-n", strconv.Itoa(num), branch)
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return out.String(), err
}
