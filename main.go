package main

import (
	"time"

	"github.com/atotto/clipboard"
	"github.com/legolord208/stdutil"
	"github.com/nsf/termbox-go"
)

type position struct {
	x int
	y int
}

const width = 29
const height = 15
const centerX = width / 2
const centerY = height / 2

var screen [][]bool
var character = position{}
var drawingStart *position
var pressedC *time.Time

var running = true
var animations = true
var disableMove bool

func main() {
	screen = makeScreen()

	err := termbox.Init()
	if err != nil {
		stdutil.PrintErr("Couldn't init termbox", err)
		return
	}
	defer termbox.Close()

	// Loops
	go func() {
		for running {
			printscreen()
			time.Sleep(time.Millisecond * 10)
		}
	}()
	initpageschedule()
	for running {
		event := termbox.PollEvent()
		handleKey(event.Key, event.Ch)
	}
}

func makeScreen() [][]bool {
	screen := make([][]bool, height)
	for i := range screen {
		screen[i] = make([]bool, width)
	}
	return screen
}

func handleKey(key termbox.Key, char rune) {
	switch key {
	case termbox.KeyArrowUp:
		move(0, -1, false)
	case termbox.KeyArrowLeft:
		move(-1, 0, false)
	case termbox.KeyArrowDown:
		move(0, 1, false)
	case termbox.KeyArrowRight:
		move(1, 0, false)

	case termbox.KeySpace:
		if disableMove {
			return
		}

		if drawingStart == nil {
			drawingStart = &position{}
			*drawingStart = character
		} else {
			start := min(*drawingStart, character)
			end := max(*drawingStart, character)

			drawingStart = nil
			fill(start, end, func(x, y int, state bool) bool {
				return !state
			})
		}

	case termbox.KeyCtrlE:
		if !clipboard.Unsupported {
			clip(getscreen())
		}
	case termbox.KeyCtrlD:
		if !clipboard.Unsupported {
			clip(getscreensquare())
		}
	case termbox.KeyCtrlL:
		if disableMove {
			return
		}
		if !clipboard.Unsupported {
			fromString(getclip())
		}

	case termbox.KeyEsc:
		running = !running
	}

	switch char {
	case 'w':
		y := 0
		if character.y > centerY {
			y = centerY
		}
		moveTo(character.x, y, false)
	case 'a':
		x := 0
		if character.x > centerX {
			x = centerX
		}
		moveTo(x, character.y, false)
	case 's':
		y := centerY
		if character.y >= centerY {
			y = height - 1
		}
		moveTo(character.x, y, false)
	case 'd':
		x := centerX
		if character.x >= centerX {
			x = width - 1
		}
		moveTo(x, character.y, false)

	case 'c':
		if disableMove {
			return
		}

		now := time.Now()
		if pressedC == nil || (*pressedC).Before(now) || (*pressedC).Equal(now) {
			then := now.Add(time.Second)
			pressedC = &then
		} else {
			fill(position{x: 0, y: 0}, position{x: width - 1, y: height - 1}, func(x, y int, state bool) bool {
				return false
			})
		}

	case '-':
		animations = !animations

	case 'q':
		running = !running

	case 'W':
		shift(0, -1, false)
	case 'A':
		shift(-1, 0, false)
	case 'S':
		shift(0, 1, false)
	case 'D':
		shift(1, 0, false)
	}
}
