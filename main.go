package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width = 80
	height = 15
)

type Universe [][]bool

// NewUniverse returns an empty universe.
func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// String return the universe as a string.
func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (width + 1) * height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'	
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

// Show clears the screen and display the universe.
func (u Universe) Show() {
	fmt.Println("\x0c", u.String())
}

// Set the state of the specific cell.
func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

// Seed random live cells into the universe.
func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

// Alive repotrs whether the specific cell is alive.
// If the coordinates are outside of the universe, they wrap around.
func (u Universe) Alive (x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

// Neighbors count the adjacent cells that are alive.
func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x + h, y + v) {
				n++
			}
		}
	}
	return n
}

// Next returns the state of the specific cell at the next step.
func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

// Step updates the state of the next universe "b" from the current universe "a".
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x:= 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()
	
	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a //Swap universes
	}
}