package main

import (
	"errors"
	"fmt"
	"task_01/internal"
)

const DEFAULT_DISCOUNT = 500

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}

	fmt.Printf("%+v\n", cust)
	fmt.Println(internal.CalcPrice(*cust, 20000))
}
