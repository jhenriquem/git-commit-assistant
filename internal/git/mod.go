package git

import (
	"bytes"
	"os/exec"
	"strings"
)

type Data struct {
	Unadded     string
	Uncommitted string
}

func Exist_repository() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return false
	}

	return strings.TrimSpace(out.String()) == "true"
}

func Get_unadded_changes() (string, error) {
	cmd := exec.Command("git", "diff")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func Get_last_commit() (string, error) {
	cmd := exec.Command("git", "log", "--oneline", "--graph")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.Split(string(output), "\n")[0], nil
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
