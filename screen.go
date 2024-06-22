package main

import (
	"fmt"
)

const (
	WIDTH  = 200
	HEIGHT = 68
)

func colorXY(screen []byte, xter byte, x, y, width, height int) []byte {
	if (x+width) < WIDTH && (y+height) < HEIGHT {
		pos := y * WIDTH
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				screen[pos+x+j] = xter
			}
			y++
			pos = (y*WIDTH)
		}
	}
	return screen
}

func Screen(space byte) []byte {
	screen := make([]byte, HEIGHT*WIDTH)

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			screen[i*WIDTH+j] = space
		}
	}
	return screen
}

func PrintScreen(screen []byte) {
	for i := 0; i < WIDTH*HEIGHT; i++ {
		if i % WIDTH == 0 {
			fmt.Printf("\n\t\t%c", screen[i])
		} else {		
			fmt.Printf("%c", screen[i])
		}
	}
	fmt.Println()
}
