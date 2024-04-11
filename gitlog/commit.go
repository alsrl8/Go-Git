package gitlog

import "strings"

type Commit struct {
	Hash    string
	Author  string
	Email   string
	Date    string
	message string
}

func ParseGitLogs(logs string) (commits []Commit) {
	if logs == "" {
		return
	}

	separator := "%%\n"
	logEntries := strings.Split(logs, separator)

	for _, logEntry := range logEntries {
		s := strings.Split(logEntry, "\n")
		message := strings.Join(s[4:], "\n")
		if strings.HasSuffix(message, "%%") {
			message = strings.TrimSuffix(message, "%%")
		}
		message = strings.TrimRight(message, "\n")

		commit := Commit{
			Hash:    strings.Trim(s[0], "\n"),
			Author:  strings.Trim(s[1], "\n"),
			Email:   strings.Trim(s[2], "\n"),
			Date:    strings.Trim(s[3], "\n"),
			message: message,
		}
		commits = append(commits, commit)
	}

	return
}
