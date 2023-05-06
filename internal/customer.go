package internal

import "errors"

const DEFAULT_DISCOUNT = 500

type Customer struct {
	Name string
	Age  int
	*Overduer
	discount bool
}

type Overduer struct {
	balance int
	debt    int
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name: name,
		Age:  age,
		Overduer: &Overduer{
			balance: balance,
			debt:    debt,
		},
		discount: discount,
	}
}

func (cust *Customer) CalcDiscount() (int, error) {
	if !cust.discount {
		return 0, errors.New("discount not available")
	}
	result := DEFAULT_DISCOUNT - cust.debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}

func (cust *Customer) CalcPrice(price int) (int, error) {
	if res, err := cust.CalcDiscount(); err != nil {
		return 0, err
	} else {
		return price - res, nil
	}
}

func (cust *Customer) WrOffDebt() error {
	if cust.debt >= cust.balance {
		return errors.New("not possible write off")
	}

	cust.balance -= cust.debt
	cust.debt = 0

	return nil
}
