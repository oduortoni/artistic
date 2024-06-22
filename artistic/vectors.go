package artistic

type Vector struct {
	X, Y, Z float64
	Direction float64
}

func Add(v1, v2 *Vector) *Vector {
	return &Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func Sub(v1, v2 *Vector) *Vector {
	return &Vector{
		X: (v1.X - v2.X) * -1,
		Y: (v1.Y - v2.Y) * -1,
		Z: (v1.Z - v2.Z) * -1,
	}
}
