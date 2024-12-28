package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	gotmux "github.com/jubnzv/go-tmux"
)

func ExecCmd(args []string) error {
	tmux, err := exec.LookPath("tmux")
	if err != nil {
		return err
	}

	args = append([]string{tmux}, args...)

	if err := syscall.Exec(tmux, args, os.Environ()); err != nil {
		return err
	}

	return nil
}

func ExecCmdWithoutStop(args []string) error {
	tmux, err := exec.LookPath("tmux")
	if err != nil {
		return fmt.Errorf("tmux not found: %w", err)
	}

	cmd := exec.Command(tmux, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running tmux: %w", err)
	}

	return nil
}

func main() {
	fzfCmd := "find ~/code ~/.config/ ~/ -mindepth 1 -maxdepth 2 -type d | fzf"
	cmd := exec.Command("bash", "-c", fzfCmd)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running fzf command: %v\n", err)
	}
	if len(string(output)) < 1 {
		os.Exit(0)	
	}
	selectedPathSlugs := strings.Split(string(output), "/")

	selectedName := strings.ReplaceAll(strings.Trim(strings.Join(selectedPathSlugs[3:], "/"), "\n"), ".", "_")
	selectedPath := strings.TrimSpace(strings.Trim(string(output), "\n"))

	server := new(gotmux.Server)
	sessions, err := server.ListSessions()
	if err != nil {
		fmt.Printf("error listing sessions: %v\n", err)
	}

	if len(sessions) == 0 && !gotmux.IsInsideTmux() {
		fmt.Println("creating new session")
		err := ExecCmd([]string{"new-session", "-s", selectedName, "-c", selectedPath})
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
		err := ExecCmdWithoutStop([]string{"new-session", "-d", "-s", selectedName, "-c", selectedPath})
		if err != nil {
			fmt.Printf("Error ExecCmd() new-session detached with name and path -> %v", err)
		}
	}
	if !hasSession && !gotmux.IsInsideTmux() {
		err := ExecCmd([]string{"new-session", "-s", selectedName, "-c", selectedPath})
		if err != nil {
			fmt.Printf("Error ExecCmd() new-session with name and path -> %v", err)
		}
	}

	if !gotmux.IsInsideTmux() {
		err := ExecCmd([]string{"a", "-t", selectedName})
		if err != nil {
			fmt.Printf("Error ExecCmd() attach with name -> %v", err)
		}
	}

	err = ExecCmd([]string{"switch-client", "-t", selectedName})
	if err != nil {
		fmt.Printf("Error ExecCmd() switch-client with name -> %v", err)
	}

}
