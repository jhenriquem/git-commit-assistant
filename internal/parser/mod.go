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

	message := strings.Split(list[0], " ")
	commit_message := strings.Join(message[2:], " ")
	return commit_message
}
