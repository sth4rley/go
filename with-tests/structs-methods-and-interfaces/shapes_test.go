package main

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("retangulo", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f\n", got, want)
		}

	})

	t.Run("circulo", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 62.83

		tolerancia := 0.01

		if got-want > tolerancia {
			t.Errorf("got %.15f want %.15f\n", got, want)
		}
	})

}

func TestArea(t *testing.T) {

	// Table driven tests
	// slice of structs, "anonymous struct"
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{Base: 12, Height: 6}, 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

	/*
		checkArea := func(t testing.TB, shape Shape, want float64) {
			t.Helper()
			got := shape.Area()
			if got != want {
				t.Errorf("got %g want %g\n", got, want)
			}
		}

		t.Run("Retangulo", func(t *testing.T) {
			rectangle := Rectangle{10.0, 10.0}
			checkArea(t, rectangle, 100.0)
		})

		t.Run("Circulo", func(t *testing.T) {
			circle := Circle{10}
			checkArea(t, circle, 314.1592653589793)
		})
	*/
}
