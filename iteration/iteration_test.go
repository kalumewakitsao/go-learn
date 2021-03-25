package iteration

import "fmt"
import "testing"

func TestIteration(t *testing.T) {

    t.Run("test iteration as part of TDD", func(t *testing.T){
        repeated := Repeat("a", 5)
        expected := "aaaaa"

        if expected != repeated {
            t.Errorf("Expected %q but got %q", expected, repeated)
        }
    })

    t.Run("test iteration as part of TDD 20 times", func(t *testing.T){
        repeated := Repeat("x", 20)
        expected := "xxxxxxxxxxxxxxxxxxxx"

        if expected != repeated {
            t.Errorf("Expected %q but got %q", expected, repeated)
        }
    })
}

func BenchmarkRepeat(b *testing.B) {
    for i:=0; i < b.N; i++ {
        // put 5 to avoid some unending loops
        Repeat("a", 5)
    }
}

func ExampleRepeat() {
    repeated := Repeat("l", 8)
    fmt.Println(repeated)
    // Output: llllllll
}
