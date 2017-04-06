package main

func min(first position, second position) position {
	var x int
	var y int

	if first.x <= second.x {
		x = first.x
	} else {
		x = second.x
	}

	if first.y <= second.y {
		y = first.y
	} else {
		y = second.y
	}
	return position{x: x, y: y}
}
func max(first position, second position) position {
	var x int
	var y int

	if first.x >= second.x {
		x = first.x
	} else {
		x = second.x
	}

	if first.y >= second.y {
		y = first.y
	} else {
		y = second.y
	}
	return position{x: x, y: y}
}
