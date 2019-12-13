package order

import "testing"

func TestAdd(t *testing.T) {
	ord := New()
	ord.Add(Order{})

	if len(ord.Orders) != 1 {
		t.Error("Order was not added")
	}
}

func TestGetAll(t *testing.T) {
	ord := New()
	ord.Add(Order{})
	res := ord.GetAll()
	if len(res) != 1 {
		t.Error("Order was not added")
	}
}
