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
	drawText(w-len(out), 0, out, white, chalk.Black)
	drawRectangle(0, 0, w, 1, blue)
	drawRectangle(0, 2, w, h-2, darkGray)
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
