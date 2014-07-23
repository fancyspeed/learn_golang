package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4
    q := &v
    p := Vertex{X: 1}
    fmt.Println(v.X)
    fmt.Println(q.X)
    fmt.Println(p.X)
}
