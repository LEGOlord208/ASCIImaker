package main

import (
	"time"

	"github.com/atotto/clipboard"
	keyboard "github.com/jteeuwen/keyboard/termbox"
	"github.com/legolord208/stdutil"
	"github.com/nsf/termbox-go"
)

type position struct {
	x int
	y int
}

const WIDTH = 29
const HEIGHT = 15
const CENTER_X = WIDTH / 2
const CENTER_Y = HEIGHT / 2

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
	kb := keyboard.New()

	// Bindings
	kb.Bind(func() { move(0, -1, false) }, "up")
	kb.Bind(func() { move(-1, 0, false) }, "left")
	kb.Bind(func() { move(0, 1, false) }, "down")
	kb.Bind(func() { move(1, 0, false) }, "right")

	kb.Bind(func() { shift(0, -1, false) }, "shift+w")
	kb.Bind(func() { shift(-1, 0, false) }, "shift+a")
	kb.Bind(func() { shift(0, 1, false) }, "shift+s")
	kb.Bind(func() { shift(1, 0, false) }, "shift+d")

	kb.Bind(func() {
		y := 0
		if character.y > CENTER_Y {
			y = CENTER_Y
		}
		moveTo(character.x, y, false)
	}, "w")
	kb.Bind(func() {
		x := 0
		if character.x > CENTER_X {
			x = CENTER_X
		}
		moveTo(x, character.y, false)
	}, "a")
	kb.Bind(func() {
		y := CENTER_Y
		if character.y >= CENTER_Y {
			y = HEIGHT - 1
		}
		moveTo(character.x, y, false)
	}, "s")
	kb.Bind(func() {
		x := CENTER_X
		if character.x >= CENTER_X {
			x = WIDTH - 1
		}
		moveTo(x, character.y, false)
	}, "d")

	kb.Bind(func() { animations = !animations }, "-")

	kb.Bind(func() {
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
	}, "space")
	kb.Bind(func() {
		if disableMove {
			return
		}

		now := time.Now()
		if pressedC == nil || (*pressedC).Before(now) || (*pressedC).Equal(now) {
			then := now.Add(time.Second)
			pressedC = &then
		} else {
			fill(position{x: 0, y: 0}, position{x: WIDTH - 1, y: HEIGHT - 1}, func(x, y int, state bool) bool {
				return false
			})
		}
	}, "c")
	kb.Bind(func() {
		if !clipboard.Unsupported {
			clip(getscreen())
		}
	}, "ctrl+e")
	kb.Bind(func() {
		if !clipboard.Unsupported {
			clip(getscreensquare())
		}
	}, "ctrl+d")
	kb.Bind(func() {
		if disableMove {
			return
		}
		if !clipboard.Unsupported {
			fromString(getclip())
		}
	}, "ctrl+l")
	kb.Bind(func() { running = false }, "q", "escape")

	// Loops
	go func() {
		for running {
			printscreen()
			time.Sleep(time.Millisecond * 10)
		}
	}()
	initpageschedule()
	for running {
		kb.Poll(termbox.PollEvent())
	}
}

func makeScreen() [][]bool {
	screen := make([][]bool, HEIGHT)
	for i := range screen {
		screen[i] = make([]bool, WIDTH)
	}
	return screen
}
