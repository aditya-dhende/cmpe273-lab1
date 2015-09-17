package Interface

import "math"

/* test */
type Shape interface {
	perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, breadth float64
}

func (circle Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func (rect Rectangle) perimeter() float64 {
	return 2 * (rect.length + rect.breadth)
}

func getPerimeter(shape Shape) float64 {
	return shape.perimeter()
}

/*func main() {

	circle := Circle{r: 10}
	rect := Rectangle{length: 10, breadth: 5}

	fmt.Printf("Perimeter of Circle: %f\n", getPerimeter(circle))
	fmt.Printf("Perimeter of Rectangle: %f\n", getPerimeter(rect))

}
*/
