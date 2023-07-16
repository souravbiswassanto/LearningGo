type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
    return 2*r.width + 2*r.height
}

type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.radius
}

func PrintShapeDetails(s Shape) {
    fmt.Println("Area:", s.Area())
    fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
    rect := Rectangle{width: 5, height: 10}
    circle := Circle{radius: 7}

    PrintShapeDetails(rect)
    PrintShapeDetails(circle)
}
