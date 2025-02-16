package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	paths := getSearchPaths()
	fzfCmd := fmt.Sprintf("find %s -mindepth 1 -maxdepth 2 -type d | fzf", paths)
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

func getSearchPaths() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("error getting executable path:", err)
		return "~/"
	}

	exeDir := filepath.Dir(exePath)
	filePath := filepath.Join(exeDir, "search_directories.txt")

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("error reading search directories", err)
	}

	text := string(data)

	lines := strings.Split(text, "\n")

	paths := strings.Join(lines, " ")

	if len(paths) > 0 {
		return paths
	} else {
		return "~/"
	}
}
