package main

import "fmt"


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

func main() {
	a := NewUniverse()
	a.Show()
}