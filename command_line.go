package termui

import (
	"fmt"
	"strings"

	"github.com/eiannone/keyboard"
)

type CommandLine struct {
	output                           []rune
	x, y, maxVisibleOutput, position int
}

func (cmd *CommandLine) WaitInput() string {
	t := NewTerminal()

	// cmd := NewCommandLine(x, prompt, placeholder)
	//	t.CommandLineOpen = true
	t.OpenCommandLine()
	done := make(chan struct{}, 1)

	var output string

	go func() {
		for {
			select {
			case keyEvent := <-t.ReceiveCommandLine():
				switch true {

				case keyEvent.Key != 0:

					switch keyEvent.Key {
					case keyboard.KeyEsc:
						ClearRow(cmd.x, 1, t.Cols())
						t.CloseCommandLine()
						// NewTerminal().CommandLineOpen = false
						// HideCursor()

						output = ""
						done <- struct{}{}
						return

					case keyboard.KeyEnter:
						ClearRow(cmd.x, 1, t.Cols())
						t.CloseCommandLine()
						// NewTerminal().CommandLineOpen = false
						// HideCursor()
						output = cmd.GetOutput()
						done <- struct{}{}
						return

					case keyboard.KeySpace:
						cmd.InsertChar(' ')

					case keyboard.KeyArrowRight:
						cmd.NextCol()

					case keyboard.KeyArrowLeft:
						cmd.PrevCol()

					case keyboard.KeyDelete:
						cmd.DeleteChar()

					case keyboard.KeyTab:
						for i := 0; i <= 3; i++ {
							cmd.InsertChar(' ')
						}

					case keyboard.KeyBackspace, keyboard.KeyBackspace2:
						cmd.Backspace()

					case keyboard.KeyHome:

						cmd.GoToLineStart()

					case keyboard.KeyEnd:
						cmd.GoToLineEnd()
					}

				case keyEvent.Char != 0:
					cmd.InsertChar(keyEvent.Char)

				}
			}
		}
	}()
	<-done
	return output
}

func NewCommandLine(x int, offsetY int, promptLine string, promptFmt func(string) string, placeholder string) *CommandLine {
	c := &CommandLine{
		x: x,
		//	y:        y,
		// position: y,
		output: []rune(placeholder),
	}

	t := NewTerminal()

	promptLine += " "

	c.y = offsetY + len([]rune(promptLine))
	c.position = c.y
	c.maxVisibleOutput = t.Cols() - offsetY - len([]rune(promptLine))

	Cell(c.x, offsetY)

	fmt.Print(promptFmt(promptLine))

	c.Print()

	c.GoToLineEnd()
	return c
}

func (c *CommandLine) GoToLineEnd() {
	c.position = c.getLastCol()

	c.printCursor()
}

func (c *CommandLine) GoToLineStart() {
	c.position = c.y

	c.printCursor()
}

func (c *CommandLine) getLastCol() int {
	return c.y + len(c.output)
}

func (c *CommandLine) goToCol(col int) {
	fmt.Printf("\033[%d;%dH", c.x, col)
}

func (c *CommandLine) printCursor() {
	c.goToCol(c.position)
	fmt.Print("\033[?25h")
}

func (c *CommandLine) NextCol() {
	if c.position < c.getLastCol() {
		c.position++
		c.printCursor()
	}
}

func (c *CommandLine) PrevCol() {
	if c.position > c.y {
		c.position--
		c.printCursor()
	}
}

func (c *CommandLine) getCurrentIndex() int {
	return c.position - c.y
}

func (c *CommandLine) InsertChar(r rune) {
	if c.position == c.maxVisibleOutput {
		return
	}

	index := c.getCurrentIndex()

	if index > len(c.output)-1 {
		c.output = append(c.output, r)
	} else {
		c.output = append(c.output[:index], append([]rune{r}, c.output[index:]...)...)
	}

	c.position++
	c.clear()
	c.Print()
}

func (c *CommandLine) DeleteChar() {
	index := c.getCurrentIndex()

	if index > len(c.output)-1 {
		return
	} else if index == len(c.output)-1 {
		c.output = c.output[:index]
	} else if index < len(c.output)-1 {
		c.output = append(c.output[:index], c.output[index+1:]...)
	}

	c.clear()
	c.Print()
}

func (c *CommandLine) Print() {
	c.goToCol(c.y)

	var visibleOutput []rune

	if c.maxVisibleOutput > len(c.output) {
		visibleOutput = c.output
	} else {
		visibleOutput = c.output[:c.maxVisibleOutput]
	}

	for _, r := range visibleOutput {
		fmt.Printf("%c", r)
	}

	c.printCursor()
}

func (c *CommandLine) Backspace() {
	index := c.getCurrentIndex()

	if index == 0 {
		return
	}

	if index > len(c.output)-1 {
		c.output = c.output[:index-1]
	} else {
		c.output = append(c.output[:index], c.output[index+1:]...)
	}
	c.position--
	c.clear()
	c.Print()
}

func (c *CommandLine) clear() {
	c.goToCol(c.y)

	fmt.Print(strings.Repeat(" ", len(c.output)+1))
	c.goToCol(c.position)
}

func (c *CommandLine) GetOutput() string {
	var stringResult string

	for _, r := range c.output {
		stringResult += string(r)
	}
	return strings.TrimSpace(stringResult)
}
