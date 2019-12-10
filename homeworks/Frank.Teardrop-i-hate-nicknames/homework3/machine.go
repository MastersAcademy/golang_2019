package main

import(
    "errors"
    "strings"
    "strconv"
)

type Machine struct {
    items []int
    sp int
}

func (m *Machine) push(x int) {
    m.items = append(m.items, x)
    m.sp += 1
}

func (m *Machine) peek() (int, error) {
    if m.sp > 0 {
        return m.items[m.sp-1], nil
    } else {
        return 0, errors.New("Machine is empty")
    }
}

func (m *Machine) pop() (int, error) {
    if m.sp > 0 {
        m.sp -= 1
        return m.items[m.sp], nil
    } else {
        return 0, errors.New("Machine is empty")
    }
}

func MakeMachine(items []int) *Machine {
    m := &Machine{items, len(items)}
    return m
}

func MakeFromString(s string) *Machine {
    itemsStr := strings.Split(s[1:len(s)-1], " ")
    items := make([]int, 0)
    for i := len(itemsStr)-1; i >= 0; i -= 1 {
        item, err := strconv.Atoi(itemsStr[i])
        if err != nil {
            panic("Malformed machine specification: " + s)
        }
        items = append(items, item)
    }
    return MakeMachine(items)
}

func (m *Machine) String() string {
    var sb strings.Builder
    sb.WriteString("[")
    for i := m.sp-1; i >= 0; i -= 1 {
        sb.WriteString(strconv.Itoa(m.items[i]))
        if (i > 0) {
            sb.WriteString(" ")
        }
    }
    sb.WriteString("]")
    return sb.String()
}
