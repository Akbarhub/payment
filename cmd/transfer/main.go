package main

import (
	"payment/pkg/bank/card"
	"fmt"
	"payment/pkg/bank/types"
)

func main() {
cards := []types.Card{
	{
		PAN:     "5000 0000 0000 0003",
		Balance: 1000,
		Active:  true,
	},
	{
		PAN:     "5000 0000 0000 0004",
		Balance: -1000,
		Active:  true,
	},
	{
		PAN:     "5000 0000 0000 0005",
		Balance: 1000,
		Active:  false,
	},
}

fmt.Println(card.PaymentSource(cards))

}