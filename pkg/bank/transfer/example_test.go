package transfer

import (
	"bank/pkg/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN: "5058270280005900",
			Balance: -56_00,
			Active: true,
			Name: "Zarplatnaya",
			Currency: types.TJS,
		},
		{
			PAN: "5058270280030403",
			Balance: 1000_00,
			Active: true,
			Name: "Sberegatelnaya",
			Currency: types.TJS,
		},
	}

	fmt.Println(PaymentSources(cards))

	// Output:
	// [{card 5058270280030403 100000}]
}
