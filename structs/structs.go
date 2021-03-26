package structs

import "math"


type Rectangle struct {
    Width float64
    Length float64
}

type Circle struct {
    Radius float64
}


func Perimeter(rectangle Rectangle) float64 {
    return 2 * (rectangle.Width + rectangle.Length)
}

func Area(rectangle Rectangle) float64 {
    return rectangle.Width * rectangle.Length
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Length
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
