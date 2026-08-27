package number

// Money is an amount in cents.
//
// Currency is held as an integer on purpose: 0.1 + 0.2 is not 0.3 in binary
// floating point, and in a value that gets spelled out in full words that error
// becomes wrong text inside a contract.
type Money int64

// Reais returns the whole units, discarding cents.
func (m Money) Reais() int64 {
	panic("TODO: implement")
}

// Cents returns the fractional part, 0 to 99.
func (m Money) Cents() int64 {
	panic("TODO: implement")
}

// String returns the amount without a currency symbol: 1.234.567,89
func (m Money) String() string {
	panic("TODO: implement")
}

// BRL returns the amount with the currency symbol: R$ 1.234.567,89
//
// A negative amount is rendered with the sign before the symbol, -R$ 10,00.
func (m Money) BRL() string {
	panic("TODO: implement")
}

// ParseMoney reads an amount written in the Brazilian convention — dot for
// thousands, comma for decimals — with or without the currency symbol.
//
// It is the inverse of String and BRL, and the half most libraries leave out.
func ParseMoney(s string) (Money, error) {
	panic("TODO: implement")
}

// FormatInt groups an integer with dots: 1234567 becomes 1.234.567
func FormatInt(n int64) string {
	panic("TODO: implement")
}

// FormatFloat renders f with the given number of decimal places, a comma for
// the decimal separator and dots for thousands.
//
// For money, prefer Money — this is for measurements and other quantities where
// binary rounding is acceptable.
func FormatFloat(f float64, decimals int) string {
	panic("TODO: implement")
}

// ParseFloat reads a number written in the Brazilian convention.
func ParseFloat(s string) (float64, error) {
	panic("TODO: implement")
}
