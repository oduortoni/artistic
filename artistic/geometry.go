package artistic

func Ray(start, direction *Vector, distance float64) *Vector {
	return Add(start, Smul(direction.Normal, distance))
}

func Plane() {

}
