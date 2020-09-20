package types

// представляет собой денежную сумму в минимальных единицах (центы, дирамы, копейки и т.д)
type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}

// Category представляет собой категорию, в которой был совершен платёж (авто, аптеки, рестораны и т.д)
type Category string

// Payment представляет с собой информацию о платеже
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}