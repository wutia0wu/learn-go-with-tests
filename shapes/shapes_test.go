package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got: '%.2f' want: '%.2f'", got, want)
		}
	})

	t.Run("Circles", func(t *testing.T) {
		circles := Circles{5}
		got := circles.Perimeter()

		want := 31.41592653589793

		if got != want {
			t.Errorf("got: '%.2f' want: '%.2f'", got, want)
		}
	})

}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("Rctangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("Cicrles", func(t *testing.T) {
		circles := Circles{10}
		checkArea(t, circles, 314.1592653589793)
	})

}
