package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"git_commit_assistant/internal/auth"
	"git_commit_assistant/internal/git"
	"git_commit_assistant/internal/handler"
	"git_commit_assistant/internal/model"
	"git_commit_assistant/internal/parser"
	"git_commit_assistant/internal/ui"
)

func main() {
	newInstace := model.Application{}
	credentials := model.CredentialsFile{}

	ui.Introduction()

	// ---
	exist, err := auth.Check_credentials_files()
	if err != nil {
		log.Println(ui.StyleError(fmt.Sprintf("\nERROR : %s", err.Error())))
	}

	if exist {
		credentials, err = auth.Get_credentials()
		if err != nil {
			log.Println(ui.StyleError("\nERROR : We were unable to access the credentials for the LLM. "))
			log.Println(ui.StyleError(fmt.Sprintf("\nERROR : %s", err.Error())))
		}
	} else {
		if err := auth.Set_credentials(&credentials); err != nil {
			log.Println(ui.StyleError(fmt.Sprintf("\nERROR : %s", err.Error())))
		} else {
			fmt.Println(ui.Bold("\n:: Alright, let's go. 👍"))
		}
	}

	// ---
	if git.Exist_repository() {

		//------
		unadded, err := git.Get_unadded_changes()
		if err != nil {
			log.Println(ui.StyleError(fmt.Sprintf("\nERROR : %s", err.Error())))
		}
		newInstace.UnAdded = unadded

		//------

		if err := check_unadded_changes(newInstace.UnAdded); err != nil {
			log.Println(ui.StyleError(fmt.Sprintf("\nERROR : %s", err.Error())))
		}

		//------
		uncommitted, err := git.Get_uncommitted_changes()
		if err != nil {
			log.Println(ui.StyleError(fmt.Sprintf("\nERROR : %s", err.Error())))
		}

		if uncommitted == "" {
			fmt.Println(ui.Bold("\n:: There are no changes to be committed."))
			return
		}

		newInstace.UnCommitted = uncommitted

		//------

		fmt.Println(ui.Bold("\n:: What did you do ? (Corrections, new features, improvements...)"))

		fmt.Print("==> ")
		prompt, err := bufio.NewReader(os.Stdout).ReadString('\n')
		if err != nil {
			log.Println(ui.StyleError(fmt.Sprintf("\nERROR : %s", err.Error())))
		}

		newInstace.Description = prompt

		// --------
		data := parser.Message(newInstace)

		stop := make(chan struct{})

		fmt.Print("\n")
		go ui.Loading(stop)

		resp, err := handler.LLM_message(data, credentials)
		if err != nil {
			close(stop)
			log.Println(ui.StyleError(fmt.Sprintf("\nERROR : %s", err.Error())))
		}

		close(stop)

		//---------

		if err := confirm_commit_message(parser.Get_commit_message(resp)); err != nil {
			log.Println(ui.StyleError(fmt.Sprintf("\nERROR : %s", err.Error())))
		}
	}
}

func check_unadded_changes(diff string) error {
	if diff != "" {
		fmt.Println(ui.Bold("\n:: I noticed there are changes outside the stage;"))
		fmt.Println(ui.Bold(":: Would you like me to add them ? [Type y for yes/anything else for no]"))

		fmt.Print("==> ")
		prompt, err := bufio.NewReader(os.Stdout).ReadString('\n')
		if err != nil {
			return err
		}

		prompt = strings.ToLower(strings.Trim(prompt, "\n"))
		if prompt == "y" {
			if err := git.Add_changes(); err != nil {
				return err
			}
		}
	}
	return nil
}

func confirm_commit_message(commit_message string) error {
	fmt.Print(ui.StyleCommit(":: Commit message : "))
	fmt.Print(commit_message + "\n")

	fmt.Println(ui.Bold(":: Did you like it ? [Type y for yes/anything else for no]\n"))

	fmt.Print("==> ")
	prompt, err := bufio.NewReader(os.Stdout).ReadString('\n')
	if err != nil {
		return err
	}

	prompt = strings.ToLower(strings.Trim(prompt, "\n"))
	if prompt == "y" {
		git.Commit(commit_message)

		last_commit, _ := git.Get_last_commit()

		fmt.Println(ui.StyleCommit("\nCommitted."))
		fmt.Println(ui.StyleHashCommit(last_commit))
	} else {
		fmt.Println("\n OK, bye.")
	}
	return nil
}
