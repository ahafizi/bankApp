package card

import (
	"fmt"
	"github.com/ahafizi/bank/pkg/bank/types"
)

func ExamplePaymentSources_allIsGood() {
	cards := []types.Card{
		{
			PAN:      "5058270280005900",
			Balance:  -56_00,
			Active:   true,
			Name:     "Zarplatnaya",
			Currency: types.TJS,
		},
		{
			PAN:      "5058270280030403",
			Balance:  1000_00,
			Active:   true,
			Name:     "Sberegatelnaya",
			Currency: types.TJS,
		},
		{
			PAN:      "5058270280030402",
			Balance:  2000_00,
			Active:   true,
			Name:     "Sberegatelnaya",
			Currency: types.TJS,
		},
		{
			PAN:      "5058270280030401",
			Balance:  1_00,
			Active:   false,
			Name:     "Sberegatelnaya",
			Currency: types.TJS,
		},
	}

	for _, payment := range PaymentSources(cards) {
		fmt.Println(payment.Number)
	}

	// Output:
	// 5058270280030403
	// 5058270280030402
}


func ExamplePaymentSources_lessZero() {
	cards := []types.Card{
		{
			PAN:      "5058270280005900",
			Balance:  -56_00,
			Active:   true,
			Name:     "Zarplatnaya",
			Currency: types.TJS,
		},
	}

	for _, payment := range PaymentSources(cards) {
		fmt.Println(payment.Number)
	}

	// Output:
	//
}

func ExamplePaymentSources_cardIsActive() {
	cards := []types.Card{
		{
			PAN:      "5058270280005900",
			Balance:  -56_00,
			Active:   false,
			Name:     "Zarplatnaya",
			Currency: types.TJS,
		},
		{
			PAN:      "5058270280005901",
			Balance:  -56_00,
			Active:   true,
			Name:     "Zarplatnaya",
			Currency: types.TJS,
		},
		{
			PAN:      "5058270280005902",
			Balance:  56_00,
			Active:   true,
			Name:     "Zarplatnaya",
			Currency: types.TJS,
		},
	}

	for _, payment := range PaymentSources(cards) {
		fmt.Println(payment.Number)
	}

	// Output:
	// 5058270280005902
}
