package Interface

import "testing"

func TestShapeInterface(t *testing.T) {

	var periCircle, periRect float64
	circle := Circle{radius: 10}
	rect := Rectangle{length: 10, breadth: 5}

	periCircle = getPerimeter(circle)
	periRect = getPerimeter(rect)

	if periCircle != 62.83185307179586 {
		t.Error("Expected 62.83185307179586, got", periCircle)
	}

	if periRect != 30 {
		t.Error("Expected 30 ,got", periRect)
	}

}



