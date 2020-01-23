package main

import(
    "strings"
    "strconv"
)

type Order []int

// create an order except one instance of the given item
// if the order has multiple instances, only one will be removed
func (o Order) except(item int) Order {
    found := false
    result := make([]int, 0)
    for _, orderItem := range o {
        if !found && item == orderItem {
            found = true
            continue
        }
        result = append(result, orderItem)
    }
    return result
}

func MakeOrder(s string) Order {
    items := make([]int, 0)
    numbers := strings.TrimSpace(s[1:len(s)-1])
    itemsStr := strings.Split(numbers, " ")
    if len(numbers) == 0 {
        return items
    }
    for _, itemStr := range itemsStr {
        item, err := strconv.Atoi(itemStr)
        if err != nil {
            panic("Malformed order specification: " + s)
        }
        items = append(items, item)
    }
    return items
}

func (o Order) String() string {
    var sb strings.Builder
    sb.WriteString("{")
    for idx, item := range o {
        sb.WriteString(strconv.Itoa(item))
        if (idx < len(o)-1) {
            sb.WriteString(" ")
        }
    }
    sb.WriteString("}")
    return sb.String()
}

func (o Order) Equals(other Order) bool {
    return strings.Compare(o.String(), other.String()) == 0
}