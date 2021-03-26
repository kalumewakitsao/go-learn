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

    checkArea := func(t testing.TB, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()

        if got != want {
            t.Errorf("Got %g but expected %g", got, want)
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
