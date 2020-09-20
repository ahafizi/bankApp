package card

import "github.com/ahafizi/bankapp/pkg/bank/types"

func PaymentSources(cards []types.Card) (paymentSource []types.PaymentSource)  {

	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}

		payment := types.PaymentSource{
			Type: "card",
			Number: string(card.PAN),
			Balance: card.Balance,
		}
		paymentSource = append(paymentSource, payment)
	}

	return paymentSource
}

func Issue(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,
	}
}