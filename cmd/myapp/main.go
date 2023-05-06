package main

import (
	"errors"
	"fmt"
	"task_01/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 0, true)

	cust.CalcDiscount()

	fmt.Printf("%+v\n", cust)
	fmt.Println(startTransactionDynamic(cust))
	fmt.Printf("%+v\n", cust)
	fmt.Println(cust.CalcPrice(20000))
}

func startTransactionDynamic(i interface{}) error {
	debtor, ok := i.(internal.Debtor)
	if !ok {
		return errors.New("incorrect type")
	}
	return debtor.WrOffDebt()
}
