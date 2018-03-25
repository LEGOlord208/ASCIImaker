package main

import (
	"time"

	"github.com/atotto/clipboard"
	"github.com/nsf/termbox-go"
)

var status string
var page1 bool

func initpageschedule() {
	go func() {
		c := time.Tick(10 * time.Second)
		for _ = range c {
			page1 = !page1
		}
	}()
}
func printscreen() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}

	var start position
	var end position

	if drawingStart != nil {
		start = min(*drawingStart, character)
		end = max(*drawingStart, character)
	}

	for y, line := range screen {
		for x, char := range line {
			c := '.'
			if drawingStart != nil && x >= start.x && x <= end.x && y >= start.y && y <= end.y {
				if char {
					c = '\''
				} else {
					c = '@'
				}
			} else if char {
				c = '#'
			} else if x == centerX {
				c = '|'
			} else if y == centerY {
				c = '-'
			}

			back := termbox.ColorDefault
			if x == character.x && y == character.y {
				back = termbox.ColorWhite
			}

			termbox.SetCell(x, y, c, termbox.ColorDefault, back)
		}
	}

	printtext(0, height+3, "Press space to start a selection, and again to end.")
	printtext(0, height+4, "Press q or escape to quit.")
	if page1 {
		printtext(0, height+5, "Press w/a/s/d to quikcly move within a larger grid.")
		printtext(0, height+6, "Press c quickly twice to clear the screen.")
	} else {
		printtext(0, height+5, "Press shift + w/a/s/d to move whole screen.")
		printtext(0, height+6, "Press minus to toggle animations.")
	}

	if clipboard.Unsupported {
		printtext(0, height+8, "Clipboard unsupported")
	} else {
		printtext(0, height+8, "Press Ctrl+E to export to clipboard")
		printtext(0, height+9, "Press Ctrl+D to export as unicode squares")
		printtext(0, height+10, "Press Ctrl+L to load/import from clipboard")
	}
	printtext(0, height+12, status)

	err = termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func printtext(x int, y int, text string) {
	for i, c := range text {
		termbox.SetCell(x+i, y, c, termbox.ColorDefault, termbox.ColorDefault)
	}
}
