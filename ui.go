package main

import (
	"fmt"

	"github.com/ttacon/chalk"
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
	// Decrease x by one to correct for alignment issues
	x -= 1

	// Save current cursor position
	fmt.Printf("\033[s")

	// Set new cursor position
	fmt.Printf("\033[%v;%vH", y, x+1)

	// Print Text
	fmt.Printf("%s%s%s", color, text, chalk.Reset)

	// Restore cursor position
	fmt.Printf("\033[u")
}

func drawDialog(title string, message string, x int, y int) {
	// Specify the width and height default values
	width := 0
	height := 6

	width = len(title) + 2

	if len(message)+2 > width {
		width = len(message) + 2
	}

	// Draw Titlebar
	drawRectangle(x, y, width, height, chalk.White.NewStyle().WithBackground(chalk.Blue))

	// Draw Shadow
	drawRectangle(x+2, y+1, width, height, chalk.Black.NewStyle().WithBackground(chalk.Black))

	// Draw Body
	drawRectangle(x, y+1, width, height-1, chalk.Black.NewStyle().WithBackground(chalk.White))

	// Draw OK button background
	drawRectangle(x+1, y+4, 4, 1, chalk.Black.NewStyle().WithBackground(chalk.Blue))

	// Draw OK button text
	drawText(x+2, y+4, "OK", chalk.White.NewStyle().WithBackground(chalk.Blue))

	// Draw window title
	drawText(x, y, title, chalk.White.NewStyle().WithBackground(chalk.Blue))

	// Draw window message
	drawText(x+1, y+2, message, chalk.Black.NewStyle().WithBackground(chalk.White))
}
