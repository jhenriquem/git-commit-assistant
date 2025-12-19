package main

import (
	"fmt"
	"log"

	git_repository "git_commit_assistant/internal/git"
	"git_commit_assistant/internal/handler"
	"git_commit_assistant/internal/model"
	"git_commit_assistant/internal/parser"
	"git_commit_assistant/internal/ui"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newInstace := model.Application{}

	fmt.Print("\n    "+ui.StyleIntroduction.Render("Hi! I'm a git commit assistant"), "\n")
	if git_repository.Exist_repository() {

		//------
		unadded, err := git_repository.Get_unadded_changes()
		if err != nil {
			log.Println(ui.StyleError.Render(fmt.Sprintf("\nERROR : %s", err.Error())))
		}
		newInstace.UnAdded = unadded

		//------
		if err := check_unadded_changes(newInstace.UnAdded); err != nil {
			log.Println(ui.StyleError.Render(fmt.Sprintf("\nERROR : %s", err.Error())))
		}
		//------
		uncommitted, err := git_repository.Get_uncommitted_changes()
		if err != nil {
			log.Println(ui.StyleError.Render(fmt.Sprintf("\nERROR : %s", err.Error())))
		}

		if uncommitted == "" {
			fmt.Println("There are no changes to be committed.")
			return
		}

		newInstace.UnCommitted = uncommitted

		//------

		prompt := ui.Input("What did you do?")
		newInstace.Description = prompt

		// --------
		data := parser.Message(newInstace)

		stop := make(chan struct{})

		go ui.Loading(stop)

		resp, err := handler.Get_commit_message(data)
		if err != nil {
			close(stop)
			log.Println(ui.StyleError.Render(fmt.Sprintf("\nERROR : %s", err.Error())))
		}

		close(stop)

		//---------

		confirm_commit_message(parser.Get_commit_message(resp.Text))
	}
}

func check_unadded_changes(diff string) error {
	if diff != "" {
		prompt := ui.Select("    I noticed there are changes outside the stage;\n    Would you like me to add them ? [Y/N]")
		option := "y"

		// Why do I only analyze the "no" field?
		// The option is automatically set to "yes," so you just need to check if it wasn't selected.

		for _, char := range prompt[3] {
			if char == 'x' {
				option = "n"
			}
		}

		if option == "y" {
			if err := git_repository.Add_changes(); err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

func confirm_commit_message(commit_message string) {
	fmt.Printf(ui.StyleCommit.Render("\n\n    Commit message :"))
	fmt.Print(commit_message + "\n")

	prompt := ui.Select("    Did you like it ?")[3]

	option := true
	for _, char := range prompt {
		if char == 'x' {
			option = false
		}
	}
	if option {
		git_repository.Commit(commit_message)
	} else {
		fmt.Println("\n    OK, bye.")
	}
}
