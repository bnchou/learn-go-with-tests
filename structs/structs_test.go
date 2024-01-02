package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Perimeter of a rectangle", func(t *testing.T) {
		got := Perimeter(Rectangle{Width: 10.0, Height: 10.0})
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("Area of a rectangle", func(t *testing.T) {
		got := Area(Rectangle{Width: 20.0, Height: 10.0})
		want := 200.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
