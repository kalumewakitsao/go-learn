// CREDIT
// https://www.youtube.com/watch?v=a4HcEsJ1hIE
// package main
package pointers

import "fmt"

func main() {
    x := 7 //value
    y := &x

    // reference: where the datatype is stored on the computer memory
    fmt.Println(x, y)

    *y = 8

    fmt.Println(x, y)
}
