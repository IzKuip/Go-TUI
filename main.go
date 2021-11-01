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

	fmt.Printf("\nTerminal Size: %vx%v\n", w, h)

	drawRectangle(20, 10, 20, 10, red)
	drawRectangle(21, 11, 20, 10, green)
	drawRectangle(22, 12, 20, 10, blue)

	fmt.Print("\nPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func clrScrn() {
	fmt.Print("\033[H\033[2J")
}

func drawRectangle(x int, y int, w int, h int, color string) {
	fmt.Printf("Printing rect: x: %v, y: %v, size: %vx%v\n", x, y, w, h)
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
