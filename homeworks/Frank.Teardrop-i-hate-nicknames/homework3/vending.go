package main

import(
    "strings"
)

type State []*Machine

// Process given order in a given state of vending machines
// if order can be satisfied by this state, return state of the machines
// after processing order and true
// If the order cannot be satisfied return nil and false
// An empty order can be satisfied trivially: by doing nothing and
// keeping current state
func ProcessOrder(order Order, state State) (State, bool) {
    if len(order) == 0 {
        return state, true
    }
    // take only first element if items in order should be taken
    // sequentially
    for _, item := range order {
        nextStates := processItem(item, state)
        for _, nextState := range nextStates {
            result, isSolution := ProcessOrder(order.except(item), nextState)
            if isSolution {
                return result, true
            }
        }
    }
    
    return nil, false
}

// Process a single order item in given state of vending machines
// Since there can be multiple machines that have this item, return
// states for all possible item retrievals
func processItem(item int, state State) []State {
    results := make([]State, 0)
    updatedMachines := takeItem(item, state)
    for _, updated := range updatedMachines {
        results = append(results, replaceMachine(updated, state))
    }
    return results
}

// Try taking an item from every machine in the state.
// Return every machine that got removed item from the top
// type []*Machine signifies that this is not a state but a collection
// of machines, even though it's the same type
func takeItem(item int, state State) []*Machine {
    result := make([]*Machine, 0)
    for _, machine := range state {
        if top, err := machine.peek(); err == nil && top == item {
            _, updated, _ := machine.pop()
            result = append(result, updated)
        }
    }
    return result
}

// Find machine in the state with the same number as m
// Return new state with machine in the state replaced
func replaceMachine(replacement *Machine, state State) State {
    result := make(State, 0)
    for _, machine := range state {
        if replacement.num == machine.num {
            result = append(result, replacement)
        } else {
            result = append(result, machine)
        }
    }
    return result
}

func MakeState(s string) State {
    machines := make(State, 0)
    strs := strings.Split(s, ", ")
    for _, machineStr := range strs {
        machines = append(machines, MakeMachineFromString(machineStr))
    }
    return machines
}

func (state State) String() string {
    machineStrs := make([]string, 0)
    for _, machine := range state {
        machineStrs = append(machineStrs, machine.String())
    }
    return strings.Join(machineStrs, ", ")
}