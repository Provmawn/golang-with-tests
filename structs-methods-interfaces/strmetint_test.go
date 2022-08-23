package strmetint

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle Perimeter", shape: Rectangle{5, 5}, want: 20},
		{name: "Circle Perimeter", shape: Circle{10}, want: 62.83185307179586},
		{name: "Triangle Perimeter", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("for %#v, got %g, want %g", tt.shape, got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle Area", shape: Rectangle{10.0, 10.0}, want: 100},
		{name: "Circle Area", shape: Circle{5.0}, want: 78.53981633974483},
		{name: "Triangle Area", shape: Triangle{12, 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("for %#v, got %g, want %g", tt.shape, got, tt.want)
			}
		})
	}
}
