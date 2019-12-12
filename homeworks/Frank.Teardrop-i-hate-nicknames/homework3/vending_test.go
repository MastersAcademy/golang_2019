package main

import(
    "testing"
)

func TestProcessItemSingleMachine(t *testing.T) {
    m1 := MakeMachineFromString("[1]")
    state := []*Machine{m1}
    // order := MakeOrder("{1}")
    nextStates := processItem(1, state)
    if len(nextStates) != 1 {
        t.Errorf("Incorrect size of resulting states: %d", len(nextStates))
    } else {
        updatedM1 := nextStates[0][0]
        if !updatedM1.isEmpty() {
            t.Errorf("machine has not been updated")
        }
    }
}

func TestProcessItemSingleMachineTwoElements(t *testing.T) {
    m1 := MakeMachineFromString("[1 2]")
    state := []*Machine{m1}
    // order := MakeOrder("{1}")
    nextStates := processItem(1, state)
    if len(nextStates) != 1 {
        t.Errorf("Incorrect size of resulting states: %d", len(nextStates))
    } else {
        updatedM1 := nextStates[0][0]
        top, _ := updatedM1.peek()
        if top != 2 {
            t.Errorf("Incorrect top element")
        }
    }
}
