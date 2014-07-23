package main

import ( 
  "fmt"
  "runtime"
)

func main() {
  println("hello", "go")
  fmt.Printf("%s\n", runtime.Version())
}
