package vectors

import (
	"fmt"
	"math"
)

type Vector struct {
	x float64
	y float64
	z float64
}

func FindCosineOfAngleBetweenVectors(v1, v2 Vector) float64 {
	var cosine float64
	var numerator = calculateSumOfThePairwiseProductOfCoordinates(v1, v2)
	var denominator = math.Sqrt((v1.x*v1.x + v1.y*v1.y + v1.z*v1.z) * (v2.x*v2.x + v2.y*v2.y + v2.z*v2.z))
	cosine = numerator / denominator
	return math.Round(cosine*10000) / 10000
}

func AreVectorsPerpendicular(v1, v2 Vector) bool {
	if calculateSumOfThePairwiseProductOfCoordinates(v1, v2) == 0 {
		return true
	}
	return false
}

func FindCrossProduct(v1, v2 Vector) Vector {
	var crossProduct Vector
	crossProduct.SetX(math.Round((v1.y*v2.z-v1.z*v2.y)*100) / 100)
	crossProduct.SetY(math.Round((v1.z*v2.x-v1.x*v2.z)*100) / 100)
	crossProduct.SetZ(math.Round((v1.x*v2.y-v1.y*v2.x)*100) / 100)
	return crossProduct
}

func FindScalarProduct(v1, v2 Vector) float64 {
	return calculateSumOfThePairwiseProductOfCoordinates(v1, v2)
}

func (v *Vector) GetLength() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func calculateSumOfThePairwiseProductOfCoordinates(v1, v2 Vector) float64 {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

func NewVector(x, y, z float64) Vector {
	var v Vector
	v.SetX(x)
	v.SetY(y)
	v.SetZ(z)
	return v
}

func (v *Vector) ToString() {
	fmt.Printf("[x = %.2f, y = %.2f, z = %.2f]\n", v.x, v.y, v.z)
}

func (v *Vector) SetX(x float64) {
	v.x = x
}

func (v *Vector) SetY(y float64) {
	v.y = y
}

func (v *Vector) SetZ(z float64) {
	v.z = z
}

func (v *Vector) GetX() float64 {
	return v.x
}

func (v *Vector) GetY() float64 {
	return v.y
}

func (v *Vector) GetZ() float64 {
	return v.z
}
