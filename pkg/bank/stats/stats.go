package stats

import (
	"github.com/ahafizi/bank/pkg/bank/types"
)

func Avg(payments []types.Payment) types.Money {
	sum := 0
	countOfPayments := len(payments)
	for _, payment := range payments {
		sum += int(payment.Amount)
	}

	return types.Money(sum / countOfPayments)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	counter := 0
	sum := 0
	for _, payment := range payments {
		if category == payment.Category {
			counter++
			sum += int(payment.Amount)
		}
	}

	return types.Money(sum/counter)
}