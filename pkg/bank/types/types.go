package types

type PaymentSource struct{
	Type string 
	Number string
	Balance Money 
}

// Payment представляет информацию о платеже.
type Payment struct{
	ID		int
	Amount 	Money
}


// Money представляет собой денежную сумму в минимальных еденицах (центы, копейки, дирамы и т.д.) 
type Money int64

// Currency представляет собой код валюты
type Currency string

// КОды валюты
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN  представляет номер карты
type PAN string

// Card представляет информацию о плятёжной карте
type Card struct {
	ID			int
	PAN			PAN
	Balance		Money // исползовали Money
	Currency	Currency
	Color		string
	Name		string
	Active		bool
	MinBalance	Money
}