package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Perimeter of a rectangle", func(t *testing.T) {
		got := Rectangle{Width: 10.0, Height: 10.0}.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Area of a rectangle", func(t *testing.T) {
		checkArea(t, Rectangle{Width: 20.0, Height: 10.0}, 200.0)
	})
	t.Run("Area of circle", func(t *testing.T) {
		checkArea(t, Circle{10}, 314.1592653589793)
	})
}
