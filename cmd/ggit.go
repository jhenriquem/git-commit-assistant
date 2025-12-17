package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	git_repository "git_commit_assistant/internal/git"
	"git_commit_assistant/internal/handler"
	"git_commit_assistant/internal/model"
	"git_commit_assistant/internal/parser"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newInstace := model.Application{}

	if git_repository.Exist_repository() {

		//------
		unadded, err := git_repository.Get_unadded_changes()
		if err != nil {
			fmt.Printf("\nERROR : %s", err.Error())
		}
		newInstace.UnAdded = unadded

		//------
		if newInstace.UnAdded != "" {

			res := ""
			fmt.Println("I noticed there are changes outside the stage;")
			fmt.Println("Would you like me to add them ? [Y/N]")

			fmt.Print(" > ")
			fmt.Scan(&res)

			if strings.ToLower(res) == "y" {
				if err := git_repository.Add_changes(); err != nil {
					fmt.Printf("\nERROR : %s", err.Error())
				}
			}

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

		fmt.Println("What did you do?")
		fmt.Print(" > ")
		fmt.Scan(&newInstace.Description)

		data := parser.Message(newInstace)

		stop := make(chan bool)

		go func(stopchan chan bool) {
			frames := []string{"|", "/", "-", "\\"}
			i := 0
			for {
				select {
				case <-stopchan:
					fmt.Print("\r                          \r")
					return
				default:
					color.RGB(192, 202, 245).Printf("\r%s", frames[i%len(frames)])
					time.Sleep(200 * time.Millisecond)
					i++
				}
			}
		}(stop)

		resp, _ := handler.Get_commit_message(data)

		stop <- true

		fmt.Println(resp.Text)
		fmt.Println(parser.Get_commit_message(resp.Text))
	}
}
