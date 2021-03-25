package arrays

import "reflect"
import "testing"

func TestArrayOps(t *testing.T) {

    t.Run("Test add an array elements fixed size", func(t *testing.T){
        numbers := []int{1, 2, 3, 4}
        added := Sum(numbers)
        expected := 10

        if expected != added {
            t.Errorf("Expected %d but got %d, array: %v", expected, added, numbers)
        }
    })

    t.Run("Test add collection elements any size", func(t *testing.T){
        numbers := [ ]int{1, 2, 3, 4, 5, 6, 7, 8}
        added := Sum(numbers)
        expected := 36

        if expected != added {
            t.Errorf("Expected %d but got %d, array: %v", expected, added, numbers)
        }
    })

    t.Run("Test add collection elements any size improved", func(t *testing.T){
        numbers_1 := []int{1, 2, 3, 4}
        numbers_2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
        added := SumAll(numbers_1, numbers_2)
        expected := []int {10, 36}

        if !reflect.DeepEqual(added, expected) {
            t.Errorf("Expected %d but got %d, arrays: %v, %v", expected, added, numbers_1, numbers_2)
        }
    })

    t.Run("Test add collection tail elements", func(t *testing.T){
        numbers_1 := []int{1, 2, 3, 4}
        numbers_2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
        added := SumAllTails(numbers_1, numbers_2)
        expected := []int {9, 35}

        if !reflect.DeepEqual(added, expected) {
            t.Errorf("Expected %d but got %d, arrays: %v, %v", expected, added, numbers_1, numbers_2)
        }
    })

    t.Run("Test add collection tail elements empty list", func(t *testing.T){
        numbers_1 := []int{}
        numbers_2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
        added := SumAllTails(numbers_1, numbers_2)
        expected := []int {0, 35}

        if !reflect.DeepEqual(added, expected) {
            t.Errorf("Expected %d but got %d, arrays: %v, %v", expected, added, numbers_1, numbers_2)
        }
    })
}
