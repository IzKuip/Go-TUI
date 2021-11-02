package main

import (
	"fmt"

	"github.com/ttacon/chalk"
	"golang.org/x/term"
)

func drawRectangle(x int, y int, w int, h int, color chalk.Style) {
	x = x - 1

	// Save current cursor position
	fmt.Printf("\033[s")

	// Set new cursor position
	fmt.Printf("\033[%v;%vH", y, x+1)

	// Print rectangle
	for i := 0; i < h; i++ {
		for ii := 0; ii < w; ii++ {
			fmt.Printf("%s %s", color, chalk.Reset)
		}

		// Set the x position of the new line to the specified x position
		if i+1 != h {
			fmt.Printf("\n\033[%vC", x)
		}
	}

	// Restore cursor position
	fmt.Printf("\033[u")
}

func drawText(x int, y int, text string, color chalk.Style) {
	x -= 1                                         // Decrease x by one to correct for alignment issues
	fmt.Printf("\033[s")                           // Save current cursor position
	fmt.Printf("\033[%v;%vH", y, x+1)              // Set new cursor position
	fmt.Printf("%s%s%s", color, text, chalk.Reset) // Print Text
	fmt.Printf("\033[u")                           // Restore cursor position
}

func drawDialog(title string, message string, x int, y int) {
	// Specify the width and height default values
	width := 0
	height := 6

	// Set the dialog width to the length of the title
	width = len(title) + 2

	// Add symbol in front of title
	title = " @ " + title

	// If the length of the message is higher, set the dialog width to it instead
	if len(message)+2 > width {
		width = len(message) + 4
	}

	drawRectangle(x, y, width, height, whiteBlueBg)      // Draw Titlebar
	drawRectangle(x+2, y+1, width, height, blackBlackBg) // Draw Shadow
	drawRectangle(x, y+1, width, height-1, blackWhiteBg) // Draw Body
	drawRectangle(x+width-6, y+4, 4, 1, blackBlueBg)     // Draw OK button background
	drawText(x+width-5, y+4, "OK", whiteBlueBg)          // Draw OK button text
	drawText(x, y, title, whiteBlueBg)                   // Draw window title
	drawText(x+2, y+2, message, blackWhiteBg)            // Draw window message
}

func fillEntireScreen(color chalk.Style) {
	w, h := getTerminalSize() // get the terminal width (w) and height (h)

	for i := 0; i < h+1; i++ {
		// For the entire height of the screen draw a rectangle with the width of the screen
		drawRectangle(0, i, w, 1, color)
	}
}

func getTerminalSize() (w int, h int) {
	width, height, err := term.GetSize(0) // use the terminal module to get the terminal size

	if err == nil {
		return width, height // if there is no error getting the terminal size, return the terminal size
	}

	return 0, 0 // if there was an error getting the terminal size, return a size of 0x0.
}

var (
	whiteYellowBg chalk.Style = chalk.White.NewStyle().WithBackground(chalk.Yellow)
	whiteBlueBg   chalk.Style = chalk.White.NewStyle().WithBackground(chalk.Blue)
	whiteBlackBg  chalk.Style = chalk.White.NewStyle().WithBackground(chalk.Black)
	blackWhiteBg  chalk.Style = chalk.Black.NewStyle().WithBackground(chalk.White)
	blackBlueBg   chalk.Style = chalk.Black.NewStyle().WithBackground(chalk.Blue)
	blackBlackBg  chalk.Style = chalk.Black.NewStyle().WithBackground(chalk.Black)
)
