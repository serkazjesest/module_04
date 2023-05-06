package internal

type Overduer struct {
	balance int
	debt    int
}

func NewOverduer(balance int, debt int) *Overduer {
	return &Overduer{
		balance: balance,
		debt:    debt,
	}
}
