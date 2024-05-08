package termui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	Reset               = "\033[0m"
	Bold                = "\033[1m"
	Italic              = "\033[3m"
	Underscore          = "\033[4m"
	Backslash           = "\u2572"
	Slash               = "\u2571"
	DoubleHorizontal    = "\u2550"
	DoubleVertical      = "\u2551"
	HorizontalSeparator = "\u2500"
	Tee                 = "\u251c"
	Corner              = "\u2514"

	BorderTopLeft        = "\u250F"
	BordertTopLeftAlt    = "\u2554"
	BorderTopRight       = "\u2513"
	BorderTopRightAlt    = "\u2557"
	BorderBottomLeft     = "\u2517"
	BorderBottomLeftAlt  = "\u255A"
	BorderBottomRight    = "\u251B"
	BorderBottomRIghtAlt = "\u255D"
	BorderVertical       = "\u2503"
	BorderVerticalAlt    = "\u2551"
	BorderHorizontal     = "\u2501"
	BorderHorizontalAlt  = "\u2550"

	CursorLeft  = "\033[D"
	CursorRight = "\033[C"
	CursorHide  = "\033[?25l"
	CursorShow  = "\033[?25h"
	CursorTop   = "\033[H"
	Block       = "\u2591"
	Space       = " "
	Line        = "\u2594"

	CarriageReturn = "\r"
	ClearLine      = "\033[2K"

	Segment = Slash + Block + Slash

	// color foreground

	Black         = "\033[30m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Magenta       = "\033[35m"
	Cyan          = "\033[36m"
	White         = "\033[37m"
	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"

	// colors background

	BgBlack         = "\033[40m"
	BgRed           = "\033[41m"
	BgGreen         = "\033[42m"
	BgYellow        = "\033[43m"
	BgBlue          = "\033[44m"
	BgMagenta       = "\033[45m"
	BgCyan          = "\033[46m"
	BgWhite         = "\033[47m"
	BgBrightBlack   = "\033[100m"
	BgBrightRed     = "\033[101m"
	BgBrightGreen   = "\033[102m"
	BgBrightYellow  = "\033[103m"
	BgBrightBlue    = "\033[104m"
	BgBrightMagenta = "\033[105m"
	BgBrightCyan    = "\033[106m"
	BgBrightWhite   = "\033[107m"
)

func HideCursor() {
	fmt.Print("\033[?25l")
}

func ShowCursor() {
	fmt.Print("\033[?25h")
}

func Cell(x, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}

func ClearRow(x, y, len int) {
	Cell(x, y)
	fmt.Print(strings.Repeat(Space, len))
	Cell(x, y)
}

func ClearScreen() {
	clearCommand := ""

	switch runtime.GOOS {
	case "linux", "darwin":
		clearCommand = "clear"
	case "windows":
		clearCommand = "cls"

	}

	cmd := exec.Command(clearCommand)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
