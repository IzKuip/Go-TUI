package main

import (
	"fmt"

	"github.com/ttacon/chalk"
)

func drawRectangle(x int, y int, w int, h int, color string) {
	x = x - 1
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

func drawText(x int, y int, text string, fg string, bg chalk.Color) {
	x = x - 1
	// Save current cursor position
	fmt.Printf("\033[s")

	// Set new cursor position
	fmt.Printf("\033[%v;%vf", y, x+1)

	// Print Text
	fmt.Printf("%s%s%s%s%s", fg, bg, text, defaultColor, black)
}

func drawDialog(title string, message string, x int, y int) {
	width := 0
	height := 4

	width = len(title) + 2
	if len(message)+2 > width {
		width = len(message) + 2
	}

	drawRectangle(x, y, width, height, blue)
	drawRectangle(x, y+1, width, height-1, lightGray)
	drawText(x, y, title, lightBlueBackground, chalk.White)
	drawText(x+1, y+2, message, whiteBackground, chalk.Black)
}
