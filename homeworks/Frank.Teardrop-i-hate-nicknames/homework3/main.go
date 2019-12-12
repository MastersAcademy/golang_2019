package main

import(
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Please provide machine state and order as arguments")
        fmt.Println("For example: \"[1 2]\" \"{1 2}")
        os.Exit(1)
    }
    run(os.Args[1], os.Args[2])    
}

func run(state, order string) {
    fmt.Printf("Ordering %s from %s\n", order, state)
    final, ok := ProcessOrder(MakeOrder(order), MakeState(state))
    if ok {
        fmt.Printf("Order completed, resulting state: %s\n", final)
    } else {
        fmt.Println("This order cannot be processed")
    }
}

func run_examples() {
    run("[1]", "{}")
    run("[1 2]", "{1 2}")
    run("[1 2]", "{1 2 3}")

    run("[1], [1 2], [1 2 3 4], [1 2 3]", "{1 2 3 4}")
    run("[2 1 1 1 1 3], [3 1 1 1 1 4 2]", "{1 1 1 1 2 3 4}")
}