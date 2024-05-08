package termui

import (
	"fmt"
)

type Section struct {
	terminal          *terminal
	Content           []string
	Width             int
	Height            int
	Top               int
	Left              int
	PaddingTop        int
	PaddingBottom     int
	PaddingLeft       int
	PaddingRight      int
	BorderTopWidth    int
	BorderLeftWidth   int
	BorderRightWidth  int
	BorderBottomWidth int
	hasBorder         bool
}

func (s *Section) Print(formatting string) {
	for x := s.Top; x < s.Top+s.Height; x++ {
		for y := s.Left; y < s.Left+s.Width; y++ {
			Cell(x, y)
			fmt.Print(Space)

		}
	}

	if s.hasBorder {
		s.PrintBorder(formatting)
	}
}

func (s *Section) SetTop(top int) {
	s.Top = top
}

func (s *Section) SetLeft(left int) {
	s.Left = left
}

func NewSection() Section {
	s := Section{terminal: NewTerminal()}
	return s
}

func (s *Section) CenterHorizontally() *Section {
	s.Left = (NewTerminal().Cols() - s.Width) / 2
	return s
}

func (s *Section) CenterVertically() *Section {
	s.Top = (NewTerminal().Rows() - s.Height) / 2
	return s
}

func (s *Section) SetPadding(Top, right, bottom, Left int) *Section {
	s.PaddingTop = Top
	s.PaddingRight = right
	s.PaddingBottom = bottom
	s.PaddingLeft = Left
	return s
}

func (s *Section) ContentStart() int {
	return s.Left + s.PaddingLeft + 1
}

func (c *Section) PrintLine(row int, line string) {
	Cell(row, c.ContentStart())
	fmt.Print(line)
}

func (s *Section) PrintContent() (row int) {
	row = s.Top + s.PaddingTop + 1

	for _, line := range s.Content {
		Cell(row, s.ContentStart())
		fmt.Print(line)

		row++
	}
	return row
}

func (s *Section) ContentFirstLine() int {
	return s.Top + s.PaddingTop + s.BorderTopWidth
}

// 3

func (s *Section) ContentWidth() int {
	return s.Width - s.PaddingLeft - s.PaddingRight - s.BorderLeftWidth - s.BorderRightWidth
}

func (s *Section) ContentLines() int {
	return s.Height - s.PaddingTop - s.PaddingBottom - s.BorderTopWidth - s.BorderBottomWidth
}

func (s *Section) OutputFirstLine() int {
	return s.Top + s.BorderTopWidth + s.PaddingTop
}

func (s *Section) OutputLastLine() int {
	return s.Top + s.Height - s.PaddingBottom - s.BorderBottomWidth
}

func (s *Section) SetHeight(h int) *Section {
	s.Height = h
	return s
}

func (s *Section) SetWidth(w int) *Section {
	s.Width = w
	return s
}

func (s *Section) SetBorder() *Section {
	s.hasBorder = true
	s.BorderBottomWidth, s.BorderLeftWidth, s.BorderRightWidth, s.BorderTopWidth = 1, 1, 1, 1

	return s
}

func (s *Section) PrintBorder(formatting string) {
	Cell(s.Top, s.Left)
	fmt.Print(formatting, BorderTopLeft, Reset)

	Cell(s.Top+s.Height-1, s.Left)
	fmt.Print(formatting, BorderBottomLeft, Reset)

	Cell(s.Top, s.Left+s.Width-1)
	fmt.Print(formatting, BorderTopRight, Reset)

	Cell(s.Top+s.Height-1, s.Left+s.Width-1)
	fmt.Print(formatting, BorderBottomRight, Reset)

	// Top horizontal border

	for y := s.Left + 1; y < s.Left+s.Width-1; y++ {
		Cell(s.Top, y)
		fmt.Print(BuildString(formatting, BorderHorizontal, Reset))

	}

	// bottom horizontal border

	for y := s.Left + 1; y < s.Left+s.Width-1; y++ {
		Cell(s.Top+s.Height-1, y)
		fmt.Print(BuildString(formatting, BorderHorizontal, Reset))

	}

	// Left vertical border

	for x := s.Top + 1; x < s.Top+s.Height-1; x++ {
		Cell(x, s.Left)
		fmt.Print(BuildString(formatting, BorderVertical, Reset))

	}

	// right vertical border

	for x := s.Top + 1; x < s.Top+s.Height-1; x++ {
		Cell(x, s.Left+s.Width-1)
		fmt.Print(BuildString(formatting, BorderVertical, Reset))

	}
}
