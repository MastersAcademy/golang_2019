package main

import(
    "fmt"
)

func main() {
    // m := &Machine{}
    // m.push(150)
    // m.push(250)
    m := MakeFromString("[5 12 34 1 3 2]")
    fmt.Println(m)
}
