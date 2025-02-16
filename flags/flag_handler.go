package flags

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/paetinspier/tmux_sessionizer/commands"
)

func Parse() {
	var help bool
	var sessionName string 
	var listDirectories bool
	var newDirectory string
	var removeDirectory string
	flag.BoolVar(&help, "h", false, "help docs")
	flag.StringVar(&sessionName, "n", "", "name new tmux session")
	flag.StringVar(&sessionName, "name", "", "name new tmux session")
	flag.BoolVar(&listDirectories, "ld", false, "list directories for fzf")
	flag.StringVar(&newDirectory, "a", "", "add directory to fzf search list")
	flag.StringVar(&removeDirectory, "r", "", "remove directory from fzf search list")

	flag.Parse()

	if help {
		commands.RunHelpDocs()
		os.Exit(0)
	}

	if len(sessionName) > 0 {
		fmt.Println(sessionName)
		n := strings.Join(flag.Args(), "-")
		commands.RunTmuxSession(n)
	}

	commands.RunTmuxSession("")
}
