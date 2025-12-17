package git_repository

import (
	"os"
	"os/exec"
)

type Data struct {
	Unadded     string
	Uncommitted string
}

func Exist_repository() bool {
	if _, err := os.ReadDir(".git"); err != nil {
		return false
	}
	return true
}

func Get_unadded_changes() (string, error) {
	cmd := exec.Command("git", "diff")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func Get_uncommitted_changes() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func Add_changes() error {
	cmd := exec.Command("git", "add", ".")
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)

	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func Add_commit_description(description string) {}
