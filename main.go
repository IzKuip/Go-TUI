package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ttacon/chalk"
)

func main() {
	clrScrn()                                                                                                   // Clear the screen
	w, h := getTerminalSize()                                                                                   // Get the terminal with (w) and height (h)
	fillEntireScreen(whiteYellowBg)                                                                             // Make the entire screen Yellow
	drawRectangle(0, 0, w, 1, whiteBlueBg)                                                                      // Draw the headerbar of the screen
	drawText(0, 0, "TechWorldInc/Go-TUI", whiteBlueBg)                                                          // Draw example text one
	out := fmt.Sprint(w) + "x" + fmt.Sprint(h)                                                                  // Create variable with terminal resolution
	drawDialog("Terminal Size", "The size of your terminal is: "+out+" characters!", 4, 5)                      // Draw example dialog one
	drawRectangle(5, 12, 10, 5, redRedBg)                                                                       // Draw example rectangle one
	drawRectangle(6, 13, 10, 5, greenGreenBg)                                                                   // Draw example rectangle two
	drawRectangle(7, 14, 10, 5, whiteBlueBg)                                                                    // Draw example rectangle three
	drawText(20, 13, "You can draw rectangles wherever you want!", blackYellowBg)                               // Draw example text one
	drawText(20, 15, "Heres how to do it:", blackYellowBg)                                                      // Draw example text two
	drawText(20, 17, "drawRectangle(x int, y int, w int, h int, color chalk.Style)", blackYellowBg)             // Draw example text three
	drawText(4, 3, "To create a dialog: drawDialog(title string, message string, x int, y int)", blackYellowBg) // Draw example text four

	fmt.Printf("\033[0;0f")                   // Set the cursor position to the top of the screen
	bufio.NewReader(os.Stdin).ReadBytes('\n') // Wait for enter to be pressed
	clrScrn()                                 // Clear the screen before ending the program
}

func clrScrn() {
	fmt.Printf("%s\033[H\033[2J", chalk.Reset)
}
