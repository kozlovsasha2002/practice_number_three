package vectors_test

import (
	"math"
	"practice_number_three/vectors"
	"testing"
)

func TestGetLength(t *testing.T) {
	t.Log("Start testing the first function.")

	t.Run("whole coordinates", func(t *testing.T) {
		var v = vectors.NewVector(18, -23, 10)

		var expectedResult float64 = math.Sqrt(float64(953))
		var realResult = v.GetLength()

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})

	t.Run("fractional coordinates", func(t *testing.T) {
		var v = vectors.NewVector(14.0, -23.0, 17.5)
		var expectedResult float64 = math.Sqrt(1031.25)

		var realResult = v.GetLength()

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})

	t.Run("zero coordinates", func(t *testing.T) {
		var v = vectors.NewVector(0, 0, 0)
		var expectedResult float64 = 0

		var realResult = v.GetLength()

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})
}

func TestFindScalarProduct(t *testing.T) {
	t.Log("Start testing the second function.")

	t.Run("integer coordinates", func(t *testing.T) {
		var v1 = vectors.NewVector(-1, 5, 4)
		var v2 = vectors.NewVector(5, -2, 0)

		var expectedResult float64 = -15
		var realResult = vectors.FindScalarProduct(v1, v2)

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})

	t.Run("fractional coordinates", func(t *testing.T) {
		var v1 = vectors.NewVector(15.0, -234.5, 111.9)
		var v2 = vectors.NewVector(51.2, 34.3, -88.0)

		var expectedResult float64 = -17122.55
		var realResult = vectors.FindScalarProduct(v1, v2)

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})

	t.Run("zero coordinates", func(t *testing.T) {
		var v1, v2 = vectors.NewVector(0, 0, 0), vectors.NewVector(0, 0, 0)

		var expectedResult float64 = 0
		var realResult = vectors.FindScalarProduct(v1, v2)

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})
}
