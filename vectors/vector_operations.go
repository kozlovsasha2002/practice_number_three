package vectors

import "fmt"

type Vector struct {
	coordinateX float64
	coordinateY float64
	coordinateZ float64
}

func NewVector(x, y, z float64) Vector {
	var v = Vector{}
	v.setX(x)
	v.setY(y)
	v.setZ(z)
	return v
}

func (v *Vector) ToString() {
	fmt.Printf("[x = %.2f, y = %.2f, z = %.2f]\n", v.coordinateX, v.coordinateY, v.coordinateZ)
}

func (v *Vector) setX(x float64) {
	v.coordinateX = x
}

func (v *Vector) setY(y float64) {
	v.coordinateY = y
}

func (v *Vector) setZ(z float64) {
	v.coordinateZ = z
}

func (v *Vector) getX() float64 {
	return v.coordinateX
}

func (v *Vector) getY() float64 {
	return v.coordinateY
}

func (v *Vector) getZ() float64 {
	return v.coordinateZ
}
