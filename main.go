package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	gotmux "github.com/jubnzv/go-tmux"
	"github.com/paetinspier/tmux_sessionizer/utils"
)

func runHelpDocs() {
	utils.DisplayHelpDocs()
}

func runTmuxSession(name string) {
	selectedName, selectedPath, err := utils.RunFzFSearch()
	if err != nil {
		fmt.Println("error with RunFzFSearch()")
		os.Exit(0)
	}
	
	if name != "" {
		selectedName = name
	}

	server := new(gotmux.Server)
	sessions, err := server.ListSessions()
	if err != nil {
		fmt.Printf("error listing sessions: %v\n", err)
	}

	if len(sessions) == 0 && !gotmux.IsInsideTmux() {
		fmt.Println("creating new session")
		err := utils.ExecTmuxCmd([]string{"new-session", "-s", selectedName, "-c", selectedPath})
		if err != nil {
			fmt.Printf("Error ExecCmd() new-session with name and path -> %v", err)
		}
		os.Exit(0)
	}
	hasSession, err := server.HasSession(selectedName)
	if err != nil {
		fmt.Printf("error with HasSession: %v", err)
	}

	if !hasSession && gotmux.IsInsideTmux() {
		err := utils.ExecTmuxCmdWithoutStop([]string{"new-session", "-d", "-s", selectedName, "-c", selectedPath})
		if err != nil {
			fmt.Printf("Error ExecCmd() new-session detached with name and path -> %v", err)
		}
	}
	if !hasSession && !gotmux.IsInsideTmux() {
		err := utils.ExecTmuxCmd([]string{"new-session", "-s", selectedName, "-c", selectedPath})
		if err != nil {
			fmt.Printf("Error ExecCmd() new-session with name and path -> %v", err)
		}
	}

	if !gotmux.IsInsideTmux() {
		err := utils.ExecTmuxCmd([]string{"a", "-t", selectedName})
		if err != nil {
			fmt.Printf("Error ExecCmd() attach with name -> %v", err)
		}
	}

	err = utils.ExecTmuxCmd([]string{"switch-client", "-t", selectedName})
	if err != nil {
		fmt.Printf("Error ExecCmd() switch-client with name -> %v", err)
	}
}

func main() {
	var help bool
	var name bool
	flag.BoolVar(&help, "h", false, "help docs")
	flag.BoolVar(&name, "n", false, "help docs")
	flag.BoolVar(&name, "name", false, "help docs")

	flag.Parse()

	if help {
		runHelpDocs()
		os.Exit(0)
	}

	if name {
		n := strings.Join(flag.Args(), "-")
		runTmuxSession(n)
	}

	runTmuxSession("")
}
