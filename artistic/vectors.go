package artistic

import (
	"errors"
	"math"
)

type Vector struct {
	X             float64
	Y             float64
	Z             float64
	Length        float64
	LengthSquared float64
	Normal        *Vector
}

type VectorOps string

func NewVector(dim ...float64) (*Vector, error) {
	n := len(dim)

	var v *Vector
	if n >= 0 && n <= 3 {
		if n == 0 {
			// only Y set to 1 => upward vector
			v = &Vector{X: 0, Y: 1, Z: 0}
		} else if n == 1 {
			// only Y set to 1 => upward vector
			v = &Vector{X: dim[0], Y: dim[0], Z: dim[0]}
		} else if n == 3 {
			// only Y set to 1 => upward vector
			v = &Vector{X: dim[0], Y: dim[1], Z: dim[2]}
		} else {
			// n == 2
			// assumed to be 2D vector
			v = &Vector{X: dim[0], Y: dim[1], Z: 0}
		}
		v.Length = Length(v)
		v.LengthSquared = LengthSquared(v)
		v.Normal = Normalize(v)
		return v, nil
	} else {
		return nil, errors.New("\n\tthe number of arguments must be between 0 and 3")
	}
}

func Clone(v *Vector) *Vector {
	return &Vector{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
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
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func Smul(v *Vector, n float64) *Vector {
	return &Vector{
		X: v.X * n,
		Y: v.Y * n,
		Z: v.Z * n,
	}
}

func Sdiv(v *Vector, n float64) *Vector {
	if n != 0 {
		return &Vector{
			X: v.X / n,
			Y: v.Y / n,
			Z: v.Z / n,
		}
	} else {
		return v
	}
}

func LengthSquared(v *Vector) float64 {
	return v.X*v.X + v.Y*v.Y + v.Z + v.Z
}

func Length(v *Vector) float64 {
	return math.Sqrt(LengthSquared(v))
}

func Normalize(v *Vector) *Vector {
	if v.X != 0 || v.Y != 0 || v.Z != 0 {
		length := Length(v)
		normal := &Vector{
			X: v.X / length,
			Y: v.Y / length,
			Z: v.Z / length,
		}
		return normal
	} else {
		return Clone(v)
	}
}

func Dot(v1, v2 *Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func Cross(v1, v2 *Vector) *Vector {
	return &Vector{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}

func Angle(v1, v2 *Vector) float64 {
	combinedLength := Length(v1) * Length(v2)
	dotProduct := Dot(v1, v2)
	return math.Acos(dotProduct / combinedLength)
}

func Angle_RadToDeg(angle float64) float64 {
	return angle * (180 / math.Pi)
}

func Andle_DegToRad(angle float64) float64 {
	return angle * (math.Pi / 180)
}
