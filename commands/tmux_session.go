package commands

import (
	"fmt"
	"os"

	gotmux "github.com/jubnzv/go-tmux"
	"github.com/paetinspier/tmux_sessionizer/utils"
)

func RunTmuxSession(name string) {
	selectedName, selectedPath, err := utils.RunFzFSearchWithSearchDirectory()
	if err != nil {
		fmt.Println("error with RunFzFSearchWithSearchPath()")
		os.Exit(0)
	}

	if name != "" {
		selectedName = name
	}

	server := new(gotmux.Server)
	sessions, err := server.ListSessions()
	if err != nil && len(sessions) > 0 {
		fmt.Printf("error listing sessions: %v\n", err)
	}

	if len(sessions) == 0 && !gotmux.IsInsideTmux() {
		fmt.Println("🚀 Launching new session")
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
