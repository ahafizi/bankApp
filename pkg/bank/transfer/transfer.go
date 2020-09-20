package transfer

import "bank/pkg/types"

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