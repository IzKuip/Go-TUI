package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ttacon/chalk"
	"golang.org/x/term"
)

func main() {
	clrScrn()

	w, h := getTerminalSize()
	drawRectangle(-1, 0, w, 1, red)
	drawRectangle(-1, 2, w, h-2, green)
	fmt.Printf("\033[0;0f")

	fmt.Print("")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func clrScrn() {
	fmt.Print("\033[H\033[2J")
}

func drawRectangle(x int, y int, w int, h int, color string) {
	// Save current cursor position
	fmt.Printf("\033[s")

	// Set new cursor position
	fmt.Printf("\033[%v;%vf", y, x+1)

	// Print rectangle
	for i := 0; i < h; i++ {
		for ii := 0; ii < w; ii++ {
			fmt.Printf("%s█%s", color, chalk.Reset)
		}

		// Set the x position of the new line to the specified x position
		fmt.Printf("\n\033[%vC", x)
	}

	// Restore cursor position
	fmt.Printf("\033[u")
}

func getTerminalSize() (w int, h int) {
	width, height, err := term.GetSize(0)
	if err == nil {
		return width, height
	}
	return 0, 0
}

const (
	black        = "\033[1;30m"
	red          = "\033[1;31m"
	green        = "\033[1;32m"
	yellow       = "\033[1;33m"
	blue         = "\033[1;34m"
	magenta      = "\033[1;35m"
	cyan         = "\033[1;36m"
	lightGray    = "\033[1;37m"
	darkGray     = "\033[1;90m"
	lightRed     = "\033[1;91m"
	lightGreen   = "\033[1;92m"
	lightYellow  = "\033[1;93m"
	lightBlue    = "\033[1;94m"
	lightMagenta = "\033[1;95m"
	lightCyan    = "\033[1;96m"
	white        = "\033[1;97m"
)
