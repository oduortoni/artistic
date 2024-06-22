package main

import (
	"log"

	"github.com/oduortoni/artistic/artistic"
)

func main() {
	screen, err := artistic.NewScreen(10, 68, '.')
	if err != nil {
		log.Fatal(err)
	}
	screen.UpdateXY('+', 20, 40, 50, 10)
	screen.Display()
}
