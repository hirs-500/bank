package types
// Money представляет собой денежную сумму в минимальных единицах (дирамы, рубль, центы, и т.д )
type Money int64
// Currency представляет собой код валюты
type Currency string

// код валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)
// PAN представяет номер карты 
type PAN string

// Card представляет информацию о платежной карте
type Card struct {
	ID     int
	PAN    PAN
	Balance Money//использовали Money 
	MinBalance Money //для минимального счета 10_000
	Currency  Currency
	Color   string
	Name   string
	Active bool
} 