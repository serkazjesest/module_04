package internal

func CalcPrice(cust *Customer, price int) (int, error) {
	if res, err := cust.CalcDiscount(); err != nil {
		return 0, err
	} else {
		return price - res, nil
	}
}
