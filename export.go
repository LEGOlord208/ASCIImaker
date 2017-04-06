package main

import (
	"github.com/atotto/clipboard"
	"strings"
)

func clip(text string) {
	err := clipboard.WriteAll(text)
	if err == nil {
		status = "Copied to clipboard!"
	} else {
		status = "Couldn't copy to clipboard: " + err.Error()
	}
}

func getscreen() string {
	var lines []string
	trimStart := -1
	trimEnd := -1

	firstNotEmpty := -1
	lastNotEmpty := -1
	for y, l := range screen {
		first := -1
		last := -1
		line := ""

		for x, char := range l {
			if char {
				if first < 0 {
					first = x
				}
				last = x

				line += "#"
			} else {
				line += " "
			}
		}

		if first >= 0 && last >= 0 {
			if firstNotEmpty < 0 {
				firstNotEmpty = y
			}
			lastNotEmpty = y
			if trimStart < 0 || first < trimStart {
				trimStart = first
			}
			if last > trimEnd {
				trimEnd = last
			}
		}
		lines = append(lines, line)
	}

	first := true
	total := ""

	if trimStart >= 0 && trimEnd >= 0 {
		for y, l := range lines {
			if y < firstNotEmpty {
				continue
			}
			if y > lastNotEmpty {
				break
			}
			l = l[trimStart : trimEnd+1]

			if first {
				first = false
			} else {
				total += "\n"
			}
			total += l
		}
	}

	return total
}
func getscreenborder() string {
	str := getscreen()
	i := strings.Index(str, "\n")
	if i < 0 {
		i = len(str) + 1
	}
	i++

	border := strings.Repeat(" ", i)
	str = border + "\n" + str + "\n" + border

	str = strings.Replace(str, "\n", " \n ", -1)
	return str
}
func getscreensquare() string {
	str := getscreenborder()
	str = strings.Replace(str, "#", "⬛", -1)
	str = strings.Replace(str, " ", "⬜", -1)

	return str
}
