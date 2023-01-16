package card

import (
	"fmt"
	"payment/pkg/bank/types"
)


func ExamplePayment (){
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

	paySources := PaymentSource(cards)

	for _, paySource := range paySources {
		fmt.Println(paySource.Number)
	}
	
	// Output:
	// 5000 0000 0000 0003
}