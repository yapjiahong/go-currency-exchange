package commona

import "strings"

func FormatAmount(amount string) string {
	parts := strings.Split(amount, ".")
	integerPart := parts[0]
	decimalPart := parts[1]

	// Add commas for thousands separator
	integerPartWithCommas := addCommas(integerPart)

	// Combine integer and decimal parts
	formattedAmount := integerPartWithCommas + "." + decimalPart

	return formattedAmount
}

func addCommas(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return addCommas(s[:n-3]) + "," + s[n-3:]
}
