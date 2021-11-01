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

	fmt.Printf("\nTerminal Size: %vx%v\n\nPrinting a rectangle at position 20x10, with size 20x10...", w, h)

	drawRectangle(20, 10, 20, 10)

	fmt.Print("\n\nPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func clrScrn() {
	fmt.Print("\033[H\033[2J")
}

func drawRectangle(x int, y int, w int, h int) {
	// Save current cursor position
	fmt.Printf("\033[s")

	// Set new cursor position
	fmt.Printf("\033[%v;%vf ", y, x)

	// Print rectangle
	for i := 0; i < h; i++ {
		for ii := 0; ii < w; ii++ {
			fmt.Printf("%s█%s", chalk.White, chalk.Reset)
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
