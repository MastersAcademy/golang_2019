package order

type Order struct {
	ID       string `json:"id"`
	Products []int  `json:"products"`
}

type AllOrders struct {
	Orders []Order
}

func New() *AllOrders {
	return &AllOrders{
		Orders: []Order{},
	}
}

func (ao *AllOrders) Add(item Order) {
	ao.Orders = append(ao.Orders, item)
}

func (ao *AllOrders) GetAll() []Order {
	return ao.Orders
}

func (ao *AllOrders) GetById(id string) Order {
	for _, ord := range ao.Orders {
		if ord.ID == id {
			return ord
		}
	}
	return Order{}
}

func (ao *AllOrders) DelById(id string) {
	for i, ord := range ao.Orders {
		if ord.ID == id {
			ao.Orders[i] = ao.Orders[len(ao.Orders)-1]
			ao.Orders[len(ao.Orders)-1] = Order{}
			ao.Orders = ao.Orders[:len(ao.Orders)-1]
		}
	}
}
