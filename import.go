package main

import (
	"github.com/atotto/clipboard"
)

func getclip() string {
	str, err := clipboard.ReadAll()
	if err != nil {
		status = "Couldn't read clipboard " + err.Error()
	}
	return str
}

func fromString(str string) {
	if str == "" {
		return
	}
	copy := makeScreen()

	x := 0
	y := 0
	for _, c := range str {
		if y >= height {
			break
		}
		if c == '\n' || x >= width {
			x = 0
			y++

			continue
		}

		active := c == '#' || c == '⬛'
		copy[y][x] = active
		x++
	}

	if animations {
		fill(position{x: 0, y: 0}, position{x: width - 1, y: height - 1}, func(x, y int, state bool) bool {
			return copy[y][x]
		})
	} else {
		screen = copy
	}
	status = "Imported!"
}
