package main

import (
	"flag"
	"os"
	"strings"

	"github.com/paetinspier/tmux_sessionizer/commands"
)

func main() {
	var help bool
	var name bool
	var listDirectories bool
	flag.BoolVar(&help, "h", false, "help docs")
	flag.BoolVar(&name, "n", false, "name new tmux session")
	flag.BoolVar(&name, "name", false, "name new tmux session")
	flag.BoolVar(&listDirectories, "listDirectories", false, "list directories for fzf")

	flag.Parse()

	if help {
		commands.RunHelpDocs()
		os.Exit(0)
	}

	if name {
		n := strings.Join(flag.Args(), "-")
		commands.RunTmuxSession(n)
	}

	commands.RunTmuxSession("")
}
