package number

// Spell writes a whole number out in words: 123 becomes "cento e vinte e três".
//
// Negative numbers are prefixed with "menos". Zero is "zero".
func Spell(n int64) string {
	panic("TODO: implement")
}

// SpellMoney writes an amount out in words, naming the currency:
// "cento e vinte e três reais e sessenta e dois centavos".
//
// An amount below one real names only the cents — 0,39 is "trinta e nove
// centavos" — and a round amount names only the units.
func SpellMoney(m Money) string {
	panic("TODO: implement")
}
