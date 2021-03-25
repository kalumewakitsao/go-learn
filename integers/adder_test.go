package integers

import "fmt"
import "testing"

func TestAdder(t *testing.T) {

    t.Run("test add 2 plus 2", func(t *testing.T){
        sum := Add(2, 2)
        expected := 4

        if sum != expected {
            t.Errorf("Expected %d but got %d", expected, sum)
        }
    })

    t.Run("test add 17 plus 18", func(t *testing.T){
        sum := Add(17, 18)
        expected := 35

        if sum != expected {
            t.Errorf("Expected %d but got %d", expected, sum)
        }
    })
}

func ExampleAdd() {
    sum := Add(12, 24)
    fmt.Println(sum)
    // Output: 36
}
