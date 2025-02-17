package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"syscall"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
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

func RunFzFSearchWithSearchDirectory() (string, string, error) {
	directories := getSearchDirectories()
	fzfCmd := fmt.Sprintf("find %s -mindepth 1 -maxdepth 2 -type d | fzf", directories)
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
	selectedDirectorySlugs := strings.Split(string(output), "/")

	name := strings.ReplaceAll(strings.Trim(strings.Join(selectedDirectorySlugs[3:], "/"), "\n"), ".", "_")
	directory := strings.TrimSpace(strings.Trim(string(output), "\n"))
	return name, directory, nil
}

func runFzFSearch() (string, error) {
	fzfCmd := "find ~/ -mindepth 1 -type d | fzf"
	cmd := exec.Command("bash", "-c", fzfCmd)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running fzf command: %v\n", err)
	}
	if len(string(output)) < 1 {
		e := fmt.Errorf("string output is zero")
		return "", e
	}

	directory := strings.TrimSpace(strings.Trim(string(output), "\n"))
	return directory, nil
}

func getSearchDirectories() string {
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

	directories := strings.Join(lines, " ")

	if len(directories) > 0 {
		return directories
	} else {
		return "~/"
	}
}

func getSearchDirectoriesFilePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	exeDir := filepath.Dir(exePath)
	filePath := filepath.Join(exeDir, "search_directories.txt")

	return filePath, nil
}

func ListSearchDirectories() {
	filePath, err := getSearchDirectoriesFilePath()
	if err != nil {
		fmt.Println("error getting search directories file path:", err)
		os.Exit(1)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("error reading search directories", err)
		os.Exit(1)
	}

	text := string(data)

	directories := strings.Split(text, "\n")
	directories = directories[:len(directories)-1]

	sort.Slice(directories, func(i, j int) bool {
		directoryOne := directories[i]
		directoryTwo := directories[j]

		directoryOneDepth := strings.Count(directoryOne, "/")
		directoryTwoDepth := strings.Count(directoryTwo, "/")

		return directoryOneDepth <= directoryTwoDepth
	})

	for _, directory := range directories {
		fmt.Println("🗀 ", directory)
	}

	os.Exit(0)
}

func AddDirectory() {
	directory, err := runFzFSearch()
	if err != nil {
		fmt.Println("error with fzf search", err)
	}
	filePath, err := getSearchDirectoriesFilePath()
	if err != nil {
		fmt.Println("error getting search directories file path:", err)
		os.Exit(1)
	}

	// read current search paths
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("error reading search directories", err)
		os.Exit(1)
	}

	text := string(data)

	directories := strings.Split(text, "\n")
	directories = directories[:len(directories)-1]

	// if directory is already listed list current directories
	for _, existingDirectory := range directories {
		if directory == existingDirectory {
			ListSearchDirectories()
		}
	}
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error opening file", err)
	}

	_, err = f.WriteString(fmt.Sprintf("%s\n", directory))
	if err != nil {
		fmt.Println("error writing string to file", err)
	}

	defer f.Close()
	// on complete return directory list
	ListSearchDirectories()
}

func RemoveDirectory() {
	filePath, err := getSearchDirectoriesFilePath()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// read current search directories
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("error reading search directories", err)
		os.Exit(1)
	}

	text := string(data)

	directories := strings.Split(text, "\n")
	directories = directories[:len(directories)-1]

	if len(directories) <= 1 {
		fmt.Println("must have at least one directories")
		ListSearchDirectories()
	}

	idx, err := fuzzyfinder.Find(
		directories,
		func(i int) string { return directories[i] },
		fuzzyfinder.WithPromptString("Select directory to remove: "),
	)
	if err != nil {
		fmt.Println("Selection canceled.")
		return
	}

	updatedDirectories := append(directories[:idx], directories[idx+1:]...)

	err = writeDirectoriesToFile(updatedDirectories, filePath)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	ListSearchDirectories()
}

func writeDirectoriesToFile(directories []string, filepath string) error {
	return os.WriteFile(filepath, []byte(strings.Join(directories, "\n")+"\n"), 0644)
}
