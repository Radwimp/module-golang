package orderbook

//Orderbook is orderbook
type Orderbook struct {
	rob []*Order //Resting for BID Orders
	roa []*Order //Resting for ASK Orders
}

//New is new
func New() *Orderbook {
	var Orderbook Orderbook
	Orderbook.rob = make([]*Order, 0)
	Orderbook.roa = make([]*Order, 0)
	return &Orderbook //{}
}

//Match is match
func (ob *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	if order.Side == 1 {
		if order.Kind == 1 {
			return ob.Bidm(order)
		}
		return ob.Bidl(order)
	}
	if order.Kind == 1 {
		return ob.Askm(order)
	}
	return ob.Askl(order)
}

//Bidm is good
func (ob *Orderbook) Bidm(order *Order) ([]*Trade, *Order) {
	var vol uint64
	trades := make([]*Trade, 0)
	for _, v := range ob.roa {
		for v.Volume > 0 {
			order.Volume--
			v.Volume--
			vol++
			if order.Volume == 0 {
				trades = append(trades, &Trade{v, order, vol, v.Price})
				return trades, nil
			}
		}
		trades = append(trades, &Trade{v, order, vol, v.Price})
		vol = 0
		ob.roa = ob.roa[1:]
	}
	ob.Addrob(order)
	return trades, order
}

//Bidl is good
func (ob *Orderbook) Bidl(order *Order) ([]*Trade, *Order) {
	var vol uint64
	trades := make([]*Trade, 0)
	for _, v := range ob.roa {
		if order.Price >= v.Price && order.Volume != 0 {
			for v.Volume > 0 {
				order.Volume--
				v.Volume--
				vol++
				if order.Volume == 0 {
					trades = append(trades, &Trade{v, order, vol, v.Price})
					if v.Volume == 0 {
						ob.roa = ob.roa[1:]
					}
					return trades, nil
				}
			}
			trades = append(trades, &Trade{v, order, vol, v.Price})
			vol = 0
			ob.roa = ob.roa[1:]
		}
	}
	ob.Addrob(order)
	return trades, nil
}

//Askm is good
func (ob *Orderbook) Askm(order *Order) ([]*Trade, *Order) {
	var vol uint64
	trades := make([]*Trade, 0)
	for _, v := range ob.rob {
		for v.Volume > 0 {
			order.Volume--
			v.Volume--
			vol++
			if order.Volume == 0 {
				trades = append(trades, &Trade{v, order, vol, v.Price})
				return trades, nil
			}
		}
		trades = append(trades, &Trade{v, order, vol, v.Price})
		vol = 0
		ob.rob = ob.rob[1:]
	}
	ob.Addroa(order)
	return trades, order
}

//Askl is good
func (ob *Orderbook) Askl(order *Order) ([]*Trade, *Order) {
	var vol uint64
	trades := make([]*Trade, 0)
	for _, v := range ob.rob {
		if order.Volume != 0 && order.Price <= v.Price {
			for v.Volume > 0 {
				order.Volume--
				v.Volume--
				vol++
				if order.Volume == 0 {
					trades = append(trades, &Trade{v, order, vol, v.Price})
					if v.Volume == 0 {
						ob.rob = ob.rob[1:]
					}
					return trades, nil
				}
			}
			trades = append(trades, &Trade{v, order, vol, v.Price})
			vol = 0
			ob.rob = ob.rob[1:]
		}
	}
	ob.Addroa(order)
	return trades, nil
}

//Addrob can do it
func (ob *Orderbook) Addrob(order *Order) {
	ob.rob = append(ob.rob, order)
	p := order.Price
	if len(ob.rob) == 2 && p > ob.rob[0].Price {
		ob.rob[0], ob.rob[1] = ob.rob[1], ob.rob[0]
		return
	}
	for l := len(ob.rob) - 1; l > 0 && ob.rob[l-1].Price < p; l-- {
		ob.rob[l] = ob.rob[l-1]
		ob.rob[l-1] = order
	}
}

//Addroa can do it
func (ob *Orderbook) Addroa(order *Order) {
	ob.roa = append(ob.roa, order)
	p := order.Price
	if len(ob.roa) == 2 && p < ob.roa[0].Price {
		ob.roa[0], ob.roa[1] = ob.roa[1], ob.roa[0]
		return
	}
	for l := len(ob.roa) - 1; l > 1 && ob.roa[l-1].Price > p; l-- {
		ob.roa[l] = ob.roa[l-1]
		ob.roa[l-1] = order
	}
}
