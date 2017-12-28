package model

import (
	"fmt"
)

func CashName(money int64) string {
	amount := ""

	val := money / 1000
	if val > 0 {
		amount += fmt.Sprintf("%dpp ", val)
		money -= val * 1000
	}

	val = money / 100
	if val > 0 {
		amount += fmt.Sprintf("%dgp ", val)
		money -= val * 100
	}

	val = money / 10
	if val > 0 {
		amount += fmt.Sprintf("%dsp ", val)
		money -= val * 10
	}
	val = money / 1
	if val > 0 {
		amount += fmt.Sprintf("%dcp ", val)
	}

	if len(amount) > 0 {
		amount = amount[0 : len(amount)-1]
	}

	return amount
}
