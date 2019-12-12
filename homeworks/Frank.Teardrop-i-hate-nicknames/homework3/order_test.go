package main

import(
    "testing"
    "fmt"
)

func TestExceptRemoves(t *testing.T) {
    order := MakeOrder("{1 2 3}")
    except := order.except(1)
    expected := MakeOrder("{2 3}")
    if !expected.Equals(except) {
        complain(t, "Did not remove correctly", except, expected)
    }
}

func TestExceptNonExistent(t *testing.T) {
    order := MakeOrder("{1 2 3}")
    except := order.except(5)
    if !except.Equals(order) {
        complain(t, "Failed when removing non-existent item", except, order)
    }
}

func TestExceptEmpty(t *testing.T) {
    order := MakeOrder("{}")
    except := order.except(1)
    if !except.Equals(order) {
        complain(t, "Failed to remove from empty order", except, order)
        t.Errorf("Failed to remove from empty order")
    }
}

func complain(t *testing.T, msg string, expected, result Order) {
    t.Errorf(fmt.Sprintf("Error: %s, expected: %s, result: %s", msg, expected, result))
}
