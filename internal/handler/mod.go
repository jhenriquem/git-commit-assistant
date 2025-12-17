package handler

import (
	"context"
	"os"

	openrouter "github.com/revrost/go-openrouter"
)

func Get_commit_message(data string) (openrouter.Content, error) {
	token := os.Getenv("OPENROUTER_TOKEN")
	model := os.Getenv("OPENROUTER_MODEL")

	client := openrouter.NewClient(
		token,
		openrouter.WithXTitle("My App"),
		openrouter.WithHTTPReferer("https://myapp.com"),
	)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openrouter.ChatCompletionRequest{
			Model: model,
			Messages: []openrouter.ChatCompletionMessage{
				openrouter.UserMessage(data),
			},
		},
	)

	response := resp.Choices[0].Message.Content

	return response, err
}
