package artistic

import (
	"errors"
	"fmt"
)

type Screen struct {
	Width  int // 200
	Height int // 68
	grid   []byte
}

func (screen *Screen) UpdateXY(xter byte, x, y, width, height int) {
	if (x+width) < screen.Width && (y+height) < screen.Height {
		pos := y * screen.Width
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				screen.grid[pos+x+j] = xter
			}
			y++
			pos = (y * screen.Width)
		}
	}
}

func (screen *Screen) Display() {
	for i := 0; i < screen.Width*screen.Height; i++ {
		if i%screen.Width == 0 {
			fmt.Printf("\n\t\t%c", screen.grid[i])
		} else {
			fmt.Printf("%c", screen.grid[i])
		}
	}
	fmt.Println()
}

func NewScreen(width, height int, space byte) (*Screen, error) {
	if height > 20 && width > 30 {
		screen := &Screen{
			Width:  width,
			Height: height,
			grid:   make([]byte, height*width),
		}

		for i := 0; i < screen.Height; i++ {
			for j := 0; j < screen.Width; j++ {
				screen.grid[i*screen.Width+j] = space
			}
		}
		return screen, nil
	}
	return nil, errors.New("\n\tthe dimensions provided are too tiny.\n\tPlease specify a dimension greater than 30X20 (width,height)")
}
