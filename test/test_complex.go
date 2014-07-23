package main

import (
    "math/cmplx"
    "fmt"
)

func main() {
    var z complex128 = cmplx.Sqrt(-5+12i)
    fmt.Println("%T(%v)\n", z, z)
}
