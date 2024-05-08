package termui

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/eiannone/keyboard"
)

type terminal struct {
	rows,
	cols int
	commandLineOpen bool
	commandLineChan chan (struct {
		Key  keyboard.Key
		Char rune
	})
}

var (
	once     sync.Once
	instance *terminal
)

func (t *terminal) IsCommandLineOpen() bool {
	return t.commandLineOpen
}

func (t *terminal) CloseCommandLine() {
	t.commandLineOpen = false
	HideCursor()
}

func (t *terminal) OpenCommandLine() {
	t.commandLineOpen = true

	ShowCursor()
}

func (t *terminal) SendToCommandLine() chan<- (struct {
	Key  keyboard.Key
	Char rune
}) {
	return t.commandLineChan
}

func (t *terminal) ReceiveCommandLine() <-chan (struct {
	Key  keyboard.Key
	Char rune
}) {
	return t.commandLineChan
}

func (t *terminal) QuitCommandLine() {
	t.commandLineChan <- struct {
		Key  keyboard.Key
		Char rune
	}{Key: keyboard.KeyEsc}
}

func NewTerminal() *terminal {
	once.Do(func() {
		instance = &terminal{}

		instance.commandLineChan = make(chan struct {
			Key  keyboard.Key
			Char rune
		}, 1)

		instance.UpdateDimensions()
	})

	return instance
}

func (t *terminal) UpdateDimensions() error {
	var cmd *exec.Cmd

	switch runtime.GOOS {

	case "windows":
		{
			cmd = exec.Command("cmd", "/c", "mode")
		}
	default:
		{
			cmd = exec.Command("stty", "size")
		}
	}
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return err
	}

	st := string(out)

	rows, cols := 0, 0

	switch runtime.GOOS {
	case "windows":

		rows, _ = strconv.Atoi(strings.Split(st, "\n")[2][8:])
		cols, _ = strconv.Atoi(strings.Split(st, "\n")[3][9:])
	default:

		split := strings.Fields(st)
		if len(split) >= 2 {
			rows, _ = strconv.Atoi(split[0])
			cols, _ = strconv.Atoi(split[1])
		}
	}

	t.rows, t.cols = rows, cols

	return nil
}

func (t *terminal) Cols() int {
	return t.cols
}

func (t *terminal) Rows() int {
	return t.rows
}
