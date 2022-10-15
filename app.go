package main

import (
	"fmt"
	"practice_number_three/vectors"
)

func main() {
	var choice int = -1

	for choice != 0 {
		printMenu()
		_, err := fmt.Scan(&choice)
		if err != nil {
			return
		}

		switch choice {
		case 1:
			var x, y, z float64
			fmt.Print("Введите координаты вектора: ")
			fmt.Scan(&x, &y, &z)
			var v1 = vectors.NewVector(x, y, z)
			fmt.Printf("Длина вектора = %f.\n", v1.GetLength())
		case 2:
			var x1, y1, z1 float64
			fmt.Print("Введите координаты первого вектора: ")
			fmt.Scan(&x1, &y1, &z1)
			var v1 = vectors.NewVector(x1, y1, z1)

			var x2, y2, z2 float64
			fmt.Print("Введите координаты второго вектора: ")
			fmt.Scan(&x2, &y2, &z2)
			var v2 = vectors.NewVector(x2, y2, z2)

			fmt.Println("Скалярное произведение векторов = ", vectors.FindScalarProduct(v1, v2))
		case 3:
			var x1, y1, z1 float64
			fmt.Print("Введите координаты первого вектора: ")
			fmt.Scan(&x1, &y1, &z1)
			var v1 = vectors.NewVector(x1, y1, z1)

			var x2, y2, z2 float64
			fmt.Print("Введите координаты второго вектора: ")
			fmt.Scan(&x2, &y2, &z2)
			var v2 = vectors.NewVector(x2, y2, z2)

			fmt.Println("Векторное произведение векторов = ", vectors.FindCrossProduct(v1, v2))
		case 4:
			var x1, y1, z1 float64
			fmt.Print("Введите координаты первого вектора: ")
			fmt.Scan(&x1, &y1, &z1)
			var v1 = vectors.NewVector(x1, y1, z1)

			var x2, y2, z2 float64
			fmt.Print("Введите координаты второго вектора: ")
			fmt.Scan(&x2, &y2, &z2)
			var v2 = vectors.NewVector(x2, y2, z2)

			fmt.Println("Косинус угла = ", vectors.FindCosineOfAngleBetweenVectors(v1, v2))
		case 5:
			var x1, y1, z1 float64
			fmt.Print("Введите координаты первого вектора: ")
			fmt.Scan(&x1, &y1, &z1)
			var v1 = vectors.NewVector(x1, y1, z1)

			var x2, y2, z2 float64
			fmt.Print("Введите координаты второго вектора: ")
			fmt.Scan(&x2, &y2, &z2)
			var v2 = vectors.NewVector(x2, y2, z2)

			fmt.Println("Вектора перпендикулярны? - ", vectors.AreVectorsPerpendicular(v1, v2))
		case 0:
			fmt.Println("Выход...")
		default:
			fmt.Println("Ошибка. Попробуйте заново.")
		}
	}

}

func printMenu() {
	fmt.Println("Меню.")
	fmt.Println("1 - вычислить длину вектора")
	fmt.Println("2 - вычислить скалярное произведение векторов")
	fmt.Println("3 - вычислить векторное произведение векторов")
	fmt.Println("4 - вычислить косинус угла между векторами")
	fmt.Println("5 - определить перпендикулярны ли вектора")
	fmt.Println("0 - выход")
	fmt.Print("Your choice: ")
}
