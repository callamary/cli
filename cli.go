package cli

import (
	"fmt"
	"os"
	"sort"
)

// Command represents a CLI command with a description.
type command struct {
	Fn          func()
	Description string
}

var (
	reset      = "\033[37m"
	info       = "\033[36m"
	warning    = "\033[33m"
	err        = "\033[31m"
	success    = "\033[32m"
	commandMap = make(map[string]command)
)

// Default command available when no commands registered
func init() {
	Command("about", about, "Show short information about Callamary CLI.")
	Command("build", build, "Compiles the CLI, incorporating any new user-defined commands.")
	Command("list", listCommands, "Outputs a list of all available commands.")
}

// Register command to be used in cli
func Command(name string, fn func(), description string) {
	commandMap[name] = command{
		Fn:          fn,
		Description: description,
	}
}

// Execute the registered command
func executeCommand(name string) {
	command, exists := commandMap[name]
	if !exists {
		fmt.Println("Unknown command:", name)
		return
	}

	command.Fn()
}

// List the available commands
func listCommands() {
	Warning("\nAvailable commands:")

	var commands []string
	maxLen := 0
	for name := range commandMap {
		if len(name) > maxLen {
			maxLen = len(name)
		}
		commands = append(commands, name)
	}

	sort.Strings(commands)

	for _, name := range commands {
		command := commandMap[name]
		fmt.Printf("- %s%-*s%s \t %s%s%s\n", success, maxLen, name, reset, info, command.Description, reset)
	}
}

func Run() {
	if len(os.Args) < 2 {
		Warning("\nNo command provided.")
		listCommands()
		return
	}

	commandName := os.Args[1]
	executeCommand(commandName)
}

// Print error message
func Error(message string, e error) {
	fmt.Println(err + fmt.Sprintf("%s Error: %s", message, err) + reset)
}

// Print warning message
func Warning(message string) {
	fmt.Println(warning + message + reset)
}

// Print success message
func Success(message string) {
	fmt.Printf("%s%s%s", success, message, reset)
}
