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
	drawText(w-len(out), 0, out, chalk.White.NewStyle().WithBackground(chalk.Black))
	//drawRectangle(0, 0, w, 1, blue)
	//drawRectangle(0, 2, w, h-2, darkGray)
	drawDialog("error", "The system cannot find the file specified", 40, 6)
	drawDialog("aboobooaboobooabooboo ", "hello world", 30, 15)

	fmt.Print("")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	main()
}

func clrScrn() {
	fmt.Printf("%s\033[H\033[2J", chalk.Reset)
}

func getTerminalSize() (w int, h int) {
	width, height, err := term.GetSize(0)
	if err == nil {
		return width, height
	}
	return 0, 0
}
