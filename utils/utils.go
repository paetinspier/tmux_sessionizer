package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func ExecTmuxCmd(args []string) error {
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

func ExecTmuxCmdWithoutStop(args []string) error {
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

func RunFzFSearch() (string, string, error) {
	fzfCmd := "find ~/code ~/.config/ ~/ -mindepth 1 -maxdepth 2 -type d | fzf"
	cmd := exec.Command("bash", "-c", fzfCmd)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running fzf command: %v\n", err)
	}
	if len(string(output)) < 1 {
		e := fmt.Errorf("string output is zero")
		return "", "", e
	}
	selectedPathSlugs := strings.Split(string(output), "/")

	name := strings.ReplaceAll(strings.Trim(strings.Join(selectedPathSlugs[3:], "/"), "\n"), ".", "_")
	path := strings.TrimSpace(strings.Trim(string(output), "\n"))
	return name, path, nil
}

func DisplayHelpDocs() {
	helpDocTitle()

	fmt.Println("    Usage: ts [option]")
	fmt.Println()
	fmt.Println("    Open a new project folder.")
	fmt.Println()
	fmt.Println("    Options:")
	fmt.Println("        -h                                 Display help docs")
	fmt.Println("        -n, -name <session name>           Specify tmux session name (default session name is created using the session path)")
	fmt.Print("\n\n\n")
}

func helpDocTitle() {
	l1 := "████████╗███╗   ███╗██╗   ██╗██╗  ██╗    ███████╗███████╗███████╗██╗ ██████╗ ███╗   ██╗██╗███████╗███████╗██████╗ "
	l2 := "╚══██╔══╝████╗ ████║██║   ██║╚██╗██╔╝    ██╔════╝██╔════╝██╔════╝██║██╔═══██╗████╗  ██║██║╚══███╔╝██╔════╝██╔══██╗"
	l3 := "   ██║   ██╔████╔██║██║   ██║ ╚███╔╝     ███████╗█████╗  ███████╗██║██║   ██║██╔██╗ ██║██║  ███╔╝ █████╗  ██████╔╝"
	l4 := "   ██║   ██║╚██╔╝██║██║   ██║ ██╔██╗     ╚════██║██╔══╝  ╚════██║██║██║   ██║██║╚██╗██║██║ ███╔╝  ██╔══╝  ██╔══██╗"
	l5 := "   ██║   ██║ ╚═╝ ██║╚██████╔╝██╔╝ ██╗    ███████║███████╗███████║██║╚██████╔╝██║ ╚████║██║███████╗███████╗██║  ██║"
	l6 := "   ╚═╝   ╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═╝    ╚══════╝╚══════╝╚══════╝╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝╚══════╝╚══════╝╚═╝  ╚═╝"

	fmt.Print("\n\n\n")
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)
	fmt.Println(l5)
	fmt.Println(l6)
	fmt.Print("\n\n\n")
}
