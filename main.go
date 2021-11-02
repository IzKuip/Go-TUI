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
	out := fmt.Sprint(w) + "x" + fmt.Sprint(h)
	drawText(w-len(out), 0, out, white, black)
	/* drawRectangle(0, 0, w, 1, red)
	drawRectangle(0, 2, w, h-2, green) */
	drawRectangle(5, 5, 50, 15, red)
	drawRectangle(6, 6, 50, 15, green)
	drawRectangle(7, 7, 50, 15, blue)
	//drawDialog("error", "The system cannot find the file specified", 40, 6)
	//drawDialog("aboobooaboobooabooboo ", "hello world", 30, 15)
	fmt.Printf("\033[0;0f")

	fmt.Print("")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	clrScrn()
}

func clrScrn() {
	fmt.Printf("%s\033[H\033[2J", chalk.Reset)
}

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

func drawText(x int, y int, text string, fg string, bg string) {
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
	height := 5

	width = len(title) + 2
	if len(message)+2 > width {
		width = len(message) + 2
	}

	drawRectangle(x, y, width, height, blue)
	drawRectangle(x, y+1, width, height-1, lightGray)
	drawText(x, y, title, lightBlueBackground, white)
	drawText(x+1, y+2, message, whiteBackground, black)
}

func getTerminalSize() (w int, h int) {
	width, height, err := term.GetSize(0)
	if err == nil {
		return width, height
	}
	return 0, 0
}

const (
	defaultColor           = "\033[1;39m"
	black                  = "\033[1;30m"
	red                    = "\033[1;31m"
	green                  = "\033[1;32m"
	yellow                 = "\033[1;33m"
	blue                   = "\033[1;34m"
	magenta                = "\033[1;35m"
	cyan                   = "\033[1;36m"
	lightGray              = "\033[1;37m"
	darkGray               = "\033[1;90m"
	lightRed               = "\033[1;91m"
	lightGreen             = "\033[1;92m"
	lightYellow            = "\033[1;93m"
	lightBlue              = "\033[1;94m"
	lightMagenta           = "\033[1;95m"
	lightCyan              = "\033[1;96m"
	white                  = "\033[1;97m"
	redBackground          = "\033[1;41m"
	greenBackground        = "\033[1;42m"
	yellowBackground       = "\033[1;43m"
	blueBackground         = "\033[1;44m"
	magentaBackground      = "\033[1;45m"
	cyanBackground         = "\033[1;46m"
	lightGrayBackground    = "\033[1;47m"
	darkGrayBackground     = "\033[1;100m"
	lightRedBackground     = "\033[1;101m"
	lightGreenBackground   = "\033[1;102m"
	lightYellowBackground  = "\033[1;103m"
	lightBlueBackground    = "\033[1;104m"
	lightMagentaBackground = "\033[1;105m"
	lightCyanBackground    = "\033[1;106m"
	whiteBackground        = "\033[1;107m"
)
