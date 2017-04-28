package main

import (
	"time"
)

func moveTo(x int, y int, override bool) {
	if !override {
		if disableMove {
			return
		}
		disableMove = true
	}
	moves := 0

	if animations {
		for character.x != x || character.y != y {
			if moves > 100 {
				break
			}
			diffX := x - character.x
			diffY := y - character.y

			absX := diffX
			if absX < 0 {
				absX = -absX
			}
			absY := diffY
			if absY < 0 {
				absY = -absY
			}

			if absX > absY {
				if diffX < 0 {
					move(-1, 0, true)
				} else {
					move(1, 0, true)
				}
			} else {
				if diffY < 0 {
					move(0, -1, true)
				} else {
					move(0, 1, true)
				}
			}
			time.Sleep(5 * time.Millisecond)

			moves++
		}
	}

	teleport(x, y)
	if !override {
		disableMove = false
	}
}
func move(x int, y int, override bool) {
	if disableMove && !override {
		return
	}
	teleport(character.x+x, character.y+y)
}
func teleport(x int, y int) {
	if x < 0 {
		x = 0
	} else if x >= width {
		x = width - 1
	}

	if y < 0 {
		y = 0
	} else if y >= height {
		y = height - 1
	}

	character.x = x
	character.y = y
}

func fill(start position, end position, state func(int, int, bool) bool) {
	go func() {
		originX := character.x
		originY := character.y
		if animations {
			disableMove = true
			moveTo(start.x, start.y, true)
		}

		for x := start.x; x <= end.x; x++ {
			for y := start.y; y <= end.y; y++ {
				old := screen[y][x]
				new := state(x, y, screen[y][x])
				screen[y][x] = new

				if animations {
					if new == old {
						continue
					}
					teleport(x, y)
					time.Sleep(time.Millisecond)
				}
			}
		}

		if animations {
			moveTo(originX, originY, true)
			disableMove = false
		}
	}()
}

func shift(shiftX int, shiftY int, override bool) {
	if disableMove && !override {
		return
	}
	copy := makeScreen()

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			shiftedX := x - shiftX
			shiftedY := y - shiftY

			tile := false
			if shiftedX >= 0 && shiftedY >= 0 && shiftedX < width && shiftedY < height {
				tile = screen[shiftedY][shiftedX]
			}

			copy[y][x] = tile
		}
	}
	screen = copy
}
