package termui

import "strings"

func swapChars(chars []rune, swaps int, swapChar rune) []rune {
	for swaps > 0 {

		chars[len(chars)-swaps] = swapChar

		swaps--
	}

	return chars
}

func TrimEnd(input string, maxLen, trimLen, swaps int, swapChar rune) string {
	toTrim := []rune(input)

	if len(toTrim) > maxLen {

		toTrim = toTrim[0:trimLen]
		toTrim = swapChars(toTrim, swaps, swapChar)

	}

	return string(toTrim)
}

func AlignCenter(maxWidth int, st, padding string) string {
	len, _ := SplitString(st)
	if maxWidth > len {
		return BuildString(strings.Repeat(padding, (maxWidth-len)/2), st)
	}

	return st
}

func AlignLeft(maxWidth int, st, padding string) string {
	len, _ := SplitString(st)

	if maxWidth > len {
		return BuildString(st, strings.Repeat(padding, maxWidth-len))
	}
	return st
}

func AlignRight(maxWidth int, st, padding string) string {
	len, _ := SplitString(st)

	if maxWidth > len {
		return BuildString(strings.Repeat(padding, maxWidth-len), st)
	}
	return st
}

func StrLen(st string) int {
	len, _ := SplitString(st)
	return len
}

func SplitString(s string) (int, []rune) {
	slice := []rune(s)

	return len(slice), slice
}

func breakString(st string, maxWidth int) []string {
	l, runeSlice := SplitString(st)
	if maxWidth > l {
		return []string{st}
	}

	parts := []string{}

	start := 0
	end := maxWidth

	for {
		if start >= l {
			break
		}

		if start+maxWidth > l {
			parts = append(parts, string(runeSlice[start:]))
			break
		}

		parts = append(parts, string(runeSlice[start:end]))
		start = end
		end = end + maxWidth

	}

	return parts
}

func BuildString(substrings ...string) string {
	var builder strings.Builder

	for _, v := range substrings {
		builder.WriteString(v)
	}

	return builder.String()
}
