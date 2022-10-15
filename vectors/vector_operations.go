package vectors

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

func FindScalarProduct(v1, v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v *Vector) GetLength() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func NewVector(x, y, z float64) Vector {
	var v Vector
	v.SetX(x)
	v.SetY(y)
	v.SetZ(z)
	return v
}

func (v *Vector) ToString() {
	fmt.Printf("[x = %.2f, y = %.2f, z = %.2f]\n", v.X, v.Y, v.Z)
}

func (v *Vector) SetX(x float64) {
	v.X = x
}

func (v *Vector) SetY(y float64) {
	v.Y = y
}

func (v *Vector) SetZ(z float64) {
	v.Z = z
}

func (v *Vector) GetX() float64 {
	return v.X
}

func (v *Vector) GetY() float64 {
	return v.Y
}

func (v *Vector) GetZ() float64 {
	return v.Z
}
