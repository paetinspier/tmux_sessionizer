package flags

import (
	"flag"
	"fmt"
	"strings"

	"github.com/paetinspier/tmux_sessionizer/commands"
	"github.com/paetinspier/tmux_sessionizer/utils"
)

func Parse() {
	var help bool
	var sessionName string 
	var listDirectories bool
	var addDirectory bool
	var removeDirectory bool

	flag.BoolVar(&help, "h", false, "help docs")
	flag.StringVar(&sessionName, "n", "", "name new tmux session")
	flag.StringVar(&sessionName, "name", "", "name new tmux session")
	flag.BoolVar(&listDirectories, "ld", false, "list directories for fzf")
	flag.BoolVar(&addDirectory, "a", false, "add directory to fzf search list")
	flag.BoolVar(&removeDirectory, "r", false, "remove directory from fzf search list")

	flag.Parse()

	if help {
		commands.RunHelpDocs()
	}

	if len(sessionName) > 0 {
		fmt.Println(sessionName)
		n := strings.Join(flag.Args(), "-")
		commands.RunTmuxSession(n)
	}

	if listDirectories {
		utils.ListSearchDirectories()
	}

	if addDirectory {
		utils.AddDirectory()
	}

	if removeDirectory {
		utils.RemoveDirectory()
	}

	commands.RunTmuxSession("")
}
