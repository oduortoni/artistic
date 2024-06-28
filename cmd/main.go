package main

import (
	"fmt"
	"log"

	"github.com/oduortoni/artistic/artistic"
)

func main() {
	screen, err := artistic.NewScreen(150, 68, '.', ClearScreen)
	if err != nil {
		log.Fatal(err)
	}
	screen.UpdateXY('+', 20, 40, 50, 10)
	screen.Display()

	v1, err := artistic.NewVector(1, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	v2, err := artistic.NewVector(0, 1, 0)
	if err != nil {
		log.Fatal(err)
	}
	rad := artistic.Angle(v1, v2)
	deg := artistic.Angle_RadToDeg(rad)
	fmt.Printf("\n\nRadian(%.2f) -> Deg(%.2f)\n", rad, deg)

	cross := artistic.Cross(v1, v2)
	fmt.Println("Cross: ", cross)

	clone := artistic.Clone(cross)
	fmt.Println(clone)
	fmt.Println(artistic.Smul(clone, 4))

	// expect to be zero => orthogonal
	dot := artistic.Dot(v1, v2)
	fmt.Println("Dot: ", dot)
}
