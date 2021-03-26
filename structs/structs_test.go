package structs

import "testing"

func TestStructs(t *testing.T) {

    t.Run("Test find perimeter of a rectangle", func(t *testing.T){
        rectangle := Rectangle{10.00, 10.00}
        got := Perimeter(rectangle)
        want := 40.00

        if got != want {
            t.Errorf("Got %.2f but expected %.2f", got, want)
        }
    })
}

func TestArea(t *testing.T) {

    areaTests := []struct {
        name string
        shape Shape
        hasArea float64
    } {
        {name: "test area of Rectangle", shape: Rectangle{Width: 10.00, Length: 20.00}, hasArea: 200.00},
        {name: "test area of Circle", shape: Circle{Radius: 10.00}, hasArea: 314.1592653589793},
        {name: "test area of Triangle", shape: Triangle{Base: 12.00, Height: 6.00}, hasArea: 36.00},
    }

    for _, tt := range areaTests {
        t.Run(tt.name, func(t *testing.T){
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %g but expected %g", tt.shape, got, tt.hasArea)
            }
        })
    }


    checkArea := func(t testing.TB, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()

        if got != want {
            t.Errorf("got %g but expected %g", got, want)
        }
    }

    t.Run("Test find area of a rectangle", func(t *testing.T){
        rectangle := Rectangle{10.00, 20.00}
        checkArea(t, rectangle, 200.00)
    })

    t.Run("Test find area of a circle", func(t *testing.T){
        circle := Circle{10.00}
        checkArea(t, circle, 314.1592653589793)
    })
}
