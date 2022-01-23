package main

import (
	"fmt"
)

// Life is an active pixel.
type Life struct {
	x int
	y int
}

func (life *Life) isNear(x int, y int) bool {
	var xOver = x + 1
	var xUnder = x - 1
	var yBefore = y - 1
	var yAfter = y + 1

	return (life.x == xOver || life.x == xUnder) && (life.y == yBefore || life.y == yAfter)
}

// GetKey returns a key that is produced by the coordinates.
func (life *Life) GetKey() string {
	return fmt.Sprintf("%d.%d", life.x, life.y)
}
