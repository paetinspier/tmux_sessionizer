package commands

import (
	"fmt"
	"os"
)

func RunHelpDocs() {
	displayHelpDocs()
	os.Exit(0)
}

func displayHelpDocs() {
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
