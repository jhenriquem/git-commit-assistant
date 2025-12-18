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

	fmt.Println("\nHi! I'm a git commit assistant\n")
	if git_repository.Exist_repository() {

		//------
		unadded, err := git_repository.Get_unadded_changes()
		if err != nil {
			fmt.Printf("\nERROR : %s", err.Error())
		}
		newInstace.UnAdded = unadded

		//------
		if newInstace.UnAdded != "" {

			fmt.Println("I noticed there are changes outside the stage;")
			prompt := ui.Select("Would you like me to add them ? [Y/N]")
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
					fmt.Printf("\nERROR : %s", err.Error())
				}
			}
			fmt.Print("\r                          \r")
		}
		//------
		uncommitted, err := git_repository.Get_uncommitted_changes()
		if err != nil {
			fmt.Printf("\nERROR : %s", err.Error())
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

		// stop := make(chan bool)

		// go ui.Loading(stop)

		resp, err := handler.Get_commit_message(data)
		if err != nil {
			// stop <- true
			log.Printf("ERROR : %s", err.Error())
		}
		// stop <- true

		fmt.Println("ddd")
		//---------

		confirm_commit_message(parser.Get_commit_message(resp.Text))
	}
}

func confirm_commit_message(commit_message string) {
	fmt.Printf("\nCommit message : %s\n", commit_message)

	prompt := ui.Select("Did you like it ?")[3]

	option := true
	for _, char := range prompt {
		if char == 'x' {
			option = false
		}
	}
	if option {
		git_repository.Commit(commit_message)
	} else {
		fmt.Println("OK, bye.")
	}
}
