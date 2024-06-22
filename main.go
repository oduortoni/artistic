package main

func main() {
	screen := Screen('.')
	screen = colorXY(screen, '+', 20, 40, 50, 10)
	PrintScreen(screen)
}
