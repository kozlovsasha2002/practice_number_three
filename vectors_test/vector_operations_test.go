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

func TestFindCrossProduct(t *testing.T) {
	t.Log("Start testing the third function.")

	t.Run("integer coordinates", func(t *testing.T) {
		var v1 = vectors.NewVector(-29, 48, 37)
		var v2 = vectors.NewVector(32, -71, -31)

		var expectedResult = vectors.NewVector(1139, 285, 523)
		var realResult = vectors.FindCrossProduct(v1, v2)

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})

	t.Run("fractional coordinates", func(t *testing.T) {
		var v1 = vectors.NewVector(123.5, -31.2, 117)
		var v2 = vectors.NewVector(-211.7, -333.3, 81)

		var expectedResult = vectors.NewVector(36468.9, -34772.4, -47767.59)
		var realResult = vectors.FindCrossProduct(v1, v2)

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})

	t.Run("zero coordinates", func(t *testing.T) {
		var v1 = vectors.NewVector(0, 0, 0)
		var v2 = vectors.NewVector(0, 0, 0)

		var expectedResult = vectors.NewVector(0, 0, 0)
		var realResult = vectors.FindCrossProduct(v1, v2)

		if realResult != expectedResult {
			t.Errorf("expected result: %f != %f real result", expectedResult, realResult)
		}
	})
}

func TestAreVectorsPerpendicular(t *testing.T) {
	t.Log("Start testing the fourth function.")

	t.Run("integer coordinates (positive)", func(t *testing.T) {
		var v1 = vectors.NewVector(1, 0, -2)
		var v2 = vectors.NewVector(2, 5, 1)

		var expectedResult = true
		var realResult = vectors.AreVectorsPerpendicular(v1, v2)

		if realResult != expectedResult {
			t.Errorf("expected result: %t != %t real result", expectedResult, realResult)
		}
	})

	t.Run("integer coordinates (negative)", func(t *testing.T) {
		var v1 = vectors.NewVector(1, 0, -2)
		var v2 = vectors.NewVector(-2, 5, 1)

		var expectedResult = false
		var realResult = vectors.AreVectorsPerpendicular(v1, v2)

		if realResult != expectedResult {
			t.Errorf("expected result: %t != %t real result", expectedResult, realResult)
		}
	})

	t.Run("fractional coordinates (positive)", func(t *testing.T) {
		var v1 = vectors.NewVector(11.5, 0, -33.3)
		var v2 = vectors.NewVector(33.3, 37.7, 11.5)

		var expectedResult = true
		var realResult = vectors.AreVectorsPerpendicular(v1, v2)

		if realResult != expectedResult {
			t.Errorf("expected result: %t != %t real result", expectedResult, realResult)
		}
	})

	t.Run("fractional coordinates (negative)", func(t *testing.T) {
		var v1 = vectors.NewVector(11.2, 12.2, 33.3)
		var v2 = vectors.NewVector(-2.2, 5.5, 7.7)

		var expectedResult = false
		var realResult = vectors.AreVectorsPerpendicular(v1, v2)

		if realResult != expectedResult {
			t.Errorf("expected result: %t != %t real result", expectedResult, realResult)
		}
	})

	t.Run("zero coordinates", func(t *testing.T) {
		var v1 = vectors.NewVector(0, 0, 0)
		var v2 = vectors.NewVector(0, 0, 0)

		var expectedResult = true
		var realResult = vectors.AreVectorsPerpendicular(v1, v2)

		if realResult != expectedResult {
			t.Errorf("expected result: %t != %t real result", expectedResult, realResult)
		}
	})
}
