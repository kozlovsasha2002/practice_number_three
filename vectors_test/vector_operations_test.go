package vectors_test

import (
	"math"
	"practice_number_three/vectors"
	"testing"
)

func TestGetLength(t *testing.T) {
	t.Run("whole numbers", func(t *testing.T) {
		var v = vectors.NewVector(18, -23, 10)
		var expectedResult float64 = math.Sqrt(float64(953))

		var realResult = v.GetLength()

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})

	t.Run("fractional numbers", func(t *testing.T) {
		var v = vectors.NewVector(14.0, -23.0, 17.5)
		var expectedResult float64 = math.Sqrt(1031.25)

		var realResult = v.GetLength()

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})

	t.Run("null vector", func(t *testing.T) {
		var v = vectors.NewVector(0, 0, 0)
		var expectedResult float64 = 0

		var realResult = v.GetLength()

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})
}
