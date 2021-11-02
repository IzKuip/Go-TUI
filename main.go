package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ttacon/chalk"
)

func main() {
	clrScrn()                                                                // Clear the screen
	w, h := getTerminalSize()                                                // Get the terminal width (w) and height (h)
	fillEntireScreen(whiteYellowBg)                                          // Make the entire screen Yellow
	drawRectangle(0, 0, w, 1, whiteBlueBg)                                   // Draw the headerbar of the screen
	out := fmt.Sprint(w) + "x" + fmt.Sprint(h)                               // Create variable with terminal resolution
	drawDialog("Terminal Size", "The size of your terminal is: "+out, 12, 6) // Draw example dialog one
	fmt.Printf("\033[0;0f")                                                  // Set the cursor position to the top of the screen
	bufio.NewReader(os.Stdin).ReadBytes('\n')                                // Wait for enter to be pressed
	clrScrn()                                                                // Clear the screen before ending the program
}

func clrScrn() {
	fmt.Printf("%s\033[H\033[2J", chalk.Reset)
}
