package stats

import (
	"github.com/ahafizi/bank/pkg/bank/types"
)

func Avg(payments []types.Payment) types.Money {
	counter := 0
	sum := 0
	for _, payment := range payments {
		counter++
		sum += int(payment.Amount)
	}

	return types.Money(sum / counter)
}