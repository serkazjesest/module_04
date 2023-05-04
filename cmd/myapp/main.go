package main

import (
	"fmt"
	"task_01/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	cust.CalcDiscount()

	fmt.Printf("%+v\n", cust)
	fmt.Println(cust.CalcPrice(20000))
}
