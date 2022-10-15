package main

import (
	"fmt"
	"practice_number_three/vectors"
)

func main() {
	fmt.Println("Program start...")

	var v = vectors.NewVector(3, 4, 5)
	v.ToString()

	fmt.Println("End of program...")
}
