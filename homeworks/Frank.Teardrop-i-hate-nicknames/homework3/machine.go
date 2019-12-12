package main

import(
    "errors"
    "strings"
    "strconv"
)

// used for identifying different machine objects that represent
// the "same" machine in different states
var nextMachineNum = 0

type Machine struct {
    items []int
    num int
}

func (m *Machine) push(x int) *Machine {
    return MakeMachine(append(m.items, x), m.num)
}

func (m *Machine) peek() (int, error) {
    if !m.isEmpty() {
        return m.items[len(m.items)-1], nil
    } else {
        return 0, errors.New("Machine is empty")
    }
}

func (m *Machine) pop() (int, *Machine, error) {
    if !m.isEmpty() {
        top := m.items[len(m.items)-1]
        nextItems := m.items[:len(m.items)-1]
        return top, MakeMachine(nextItems, m.num), nil
    } else {
        return 0, nil, errors.New("Machine is empty")
    }
}

func (m *Machine) isEmpty() bool {
    return len(m.items) == 0
}

func MakeMachine(items []int, num int) *Machine {
    tmp := make([]int, len(items))
    copy(tmp, items)
    m := &Machine{tmp, num}
    return m
}

func MakeMachineFromString(s string) *Machine {
    itemsStr := strings.Split(s[1:len(s)-1], " ")
    items := make([]int, 0)
    for i := len(itemsStr)-1; i >= 0; i -= 1 {
        item, err := strconv.Atoi(itemsStr[i])
        if err != nil {
            panic("Malformed machine specification: " + s)
        }
        items = append(items, item)
    }
    nextMachineNum += 1
    return MakeMachine(items, nextMachineNum)
}

func (m *Machine) String() string {
    var sb strings.Builder
    sb.WriteString("[")
    for i := len(m.items)-1; i >= 0; i -= 1 {
        sb.WriteString(strconv.Itoa(m.items[i]))
        if (i > 0) {
            sb.WriteString(" ")
        }
    }
    sb.WriteString("]")
    return sb.String()
}
