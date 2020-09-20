package transfer

const bonusPercent = 0.005

func Bonus(amount int) (bonus int) {
	bonus = int(float64(amount) * bonusPercent)
	return bonus
}

func Total(amount int) (totalWithBonus int) {

	totalWithBonus = amount + Bonus(amount)

	return totalWithBonus
}
