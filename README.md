![cli image](./assets/screenshot.png) 
# Git Commit Assistant

A CLI assistant that automatically generates commit messages based on the diff,
the type of change, and the patch size. Additionally, the tool relies on a brief
description of the commit's purpose provided by the user. I developed it for personal use, but decided to publish it on GitHub.

### Features

- Analysis of diffs, patches, and change types
- Interpretation of user messages
- Automatic generation of commit messages
- Automatic commit (with user approval)

### Project Structure
```bash
git_assistant/
├── cmd/
    └── main.go # Entry point
├── internal/
    ├── git/     # Git functionalities
    ├── parser/  # Message parser 
    ├── handler/ # LLM request
    ├── ui/      # Interface elements
    └── model/   # Data model 
├── go.mod
└── README.md
```

### Requeriments
- Go 
- [OpenRouter account](https://openrouter.ai/) 
- OpenRouter key

 
