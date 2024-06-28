package artistic

type Point struct {
	Vector
	Figure byte
}

type Rectangle struct {
	Point
	Width, Height int
}

type Triangle struct {
	Point	
	Base, Height int
}

type Circle struct {
	Point	
	Radius int
}

type Path struct {
	Points []Point
}
