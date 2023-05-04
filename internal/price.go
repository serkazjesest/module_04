package internal

func CalcPrice(disc Discounter, price int) (int, error) {
	if res, err := disc.CalcDiscount(); err != nil {
		return 0, err
	} else {
		return price - res, nil
	}
}
