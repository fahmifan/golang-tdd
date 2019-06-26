package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerim := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		want := 40.0

		checkPerim(t, rect, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 2 * 10.0 * math.Pi

		checkPerim(t, circle, want)
	})

}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, math.Pi * 10 * 10},
	}

	for _, tt := range areaTests {
		got := tt.shape
		want := tt.want
		checkArea(t, got, want)
	}
}
