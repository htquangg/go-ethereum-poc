package visualize

import (
	"strconv"
)

func PrintAllSecretDetails(balances map[string]float64) {
	rows := [][2]string{}
	for addr, balance := range balances {
		rows = append(rows, [...]string{addr, strconv.FormatFloat(balance, 'g', 18, 64)})
	}

	headers := [...]string{"ADDRESS", "BALANCE"}

	Table(headers, rows)
}
