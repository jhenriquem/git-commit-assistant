package parser

import (
	"fmt"
	"strings"

	"git_commit_assistant/internal/model"
)

func Message(data model.Application) string {
	command := []string{
		"Generate a commit message using this data: description, diff, and diff --cached.",
		"Your response should follow this structure:",
		"Commit message: ...",
		"Summary...",
		"Changes...",
		"Implementation Details...",
		"",
		"Be direct and assertive.",
		"------------------------------------",
		"",
		"",
		"--- DESCRIPTION --- ",
		fmt.Sprintf("%s", data.Description),
		"", "",
		"--- DIFF ---  ",
		fmt.Sprintf("%s", data.UnAdded),
		"", "",
		"--- DIFF CACHED ---  ",
		fmt.Sprintf("%s", data.UnCommitted),
	}

	response := strings.Join(command, "\n")
	return response
}

func Get_commit_message(data string) string {
	list := strings.Split(data, "\n")

	commit_message := list[0][19:]
	return commit_message
}
