package main

import (
	"errors"
	"fmt"
	"task_01/internal"
)

func main() {
	overduer := internal.NewOverduer(2000, 0)
	cust := internal.NewCustomer("Dmitry", 23, *overduer, true)

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
