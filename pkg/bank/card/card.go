package card

import (
	"payment/pkg/bank/types"
)


func PaymentSource(cards []types.Card) []types.PaymentSource{

	paySource := make([]types.PaymentSource, 0, len(cards))

	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}
		
		paySource = append(paySource, types.PaymentSource{
			Type: "card",
			Number: string(card.PAN),
			Balance: types.Money(card.Balance),
		})
	}

	return paySource
}