![cli image](./assets/screenshot.png)
<h2 align="center">🤖 Git Commit Assistant</h2>

![Version](https://img.shields.io/github/v/release/henriquemco/git-commit-assistant?label=vers%C3%A3o)
![Go Version](https://img.shields.io/github/go-mod/go-version/henriquemco/git-commit-assistant)
![Made with Go](https://img.shields.io/badge/feito%20com-Go-00ADD8?logo=go)

A CLI assistant that automatically generates commit messages based on the diff,
the type of change, and the patch size. Additionally, the tool relies on a brief
description of the commit's purpose provided by the user. I developed it for personal use, but decided to publish it on GitHub.

### 📌 Features

- Analysis of diffs, patches, and change types
- Interpretation of user messages
- Automatic generation of commit messages
- Automatic commit (with user approval)

### 💡 Motivation

Writing good commit messages is important, but repetitive.
This tool aims to reduce cognitive overhead while still preserving semantic and contextual quality in commits.

### ⚠️ Disclaimer

- This tool does not push commits automatically.
- Always review generated messages before confirming.
- The quality of output depends on the chosen LLM model.

### 🗃️ Project Structure
```bash
git_assistant/
├── cmd/
    └── main.go # Entry point
├── internal/
│   ├── git/         # Git-related operations
│   ├── parser/      # User message parsing
│   ├── handler/     # LLM requests and responses
│   ├── ui/          # CLI interface and prompts
│   ├── auth/        # LLM credentials handling
│   └── model/       # Core data models
├── go.mod
└── README.md
```


### ⚙️ Requirements
Before installing and running the project, make sure you have:

- Go (version 1.20 or higher recommended)
- An OpenRouter account
    👉 [OpenRouter account](https://openrouter.ai/) 
- An OpenRouter API key


### 🚀 Installation 

 **On first run, the tool will prompt you for:** 
 - Your OpenRouter API key
 - The LLM model you want to use
 > These values are stored locally for future executions.



#### Release (Recommended)
Download a prebuilt binary for your platform from the GitHub Releases page:

👉 [Release](https://github.com/henriquemco/git-commit-assistant/releases/tag/v0.1.0)

#### Build from Source

- **Clone the repository:**

```bash 
git clone https://github.com/your-username/git-commit-assistant.git
cd git-commit-assistant
```

- **Build the project:**

```bash 
go build -o git-commit-assistant ./cmd
 ```

 - **Run the application**

```bash
 ./git-commit-assistant
 ```

> 💡 Tip: Move the binary to a directory in your $PATH (e.g. ~/.local/bin) for global access.


### 📍 Usage

Run the tool inside a Git repository with staged or unstaged changes:

```bash 
git-commit-assistant
```

**You will be prompted to:**
1. Briefly describe the purpose of the commit
2. Review the generated commit message
3. Confirm or reject the commit
No commit is created without explicit user approval.

### Configuration

Configuration is handled interactively on first use.
If needed, you can reset your credentials by deleting the local config file (location depends on OS).


