package main

import (
	"fmt"
	"math"
	"practice_number_three/vectors"
)

func main() {
	fmt.Println("Program start...")

	var v1 = vectors.NewVector(1, 6, -2)
	var v2 = vectors.NewVector(2, 5, 3)
	var len1 = v1.GetX()*v1.GetX() + v1.GetY()*v1.GetY() + v1.GetZ()*v1.GetZ()
	var len2 = v2.GetX()*v2.GetX() + v2.GetY()*v2.GetY() + v2.GetZ()*v2.GetZ()
	fmt.Println(len1)
	fmt.Println(len2)
	vectors.FindCosineOfAngleBetweenVectors(v1, v2)
	fmt.Println(math.Sqrt(len1))
	fmt.Println(math.Sqrt(len2))
	fmt.Println(math.Sqrt(len1) * math.Sqrt(len2))

	fmt.Println("End of program...")
}
