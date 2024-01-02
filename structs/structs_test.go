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
	t.Run("Area of a rectangle", func(t *testing.T) {
		got := Rectangle{Width: 20.0, Height: 10.0}.Area()
		want := 200.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("Area of circle", func(t *testing.T) {
		got := Circle{10}.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
