package arrays

import "testing"

func TestArrayOps(t *testing.T) {

    t.Run("Test add an array elements", func(t *testing.T){
        numbers := [4]int{1, 2, 3, 4}
        added := Sum(numbers)
        expected := 10

        if expected != added {
            t.Errorf("Expected %d but got %d, array: %v", expected, added, numbers)
        }
    })
}
