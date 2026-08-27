package number_test

import (
	"testing"

	"github.com/ismaelmdesousa/brdatatypes/number"
)

func TestSpellMoney(t *testing.T) {
	tests := []struct {
		in   number.Money // cents
		want string
	}{
		// Verified results carried over from the NumberToExtensive README.
		{40000, "quatrocentos reais"},
		{39, "trinta e nove centavos"},
		{129, "um real e vinte e nove centavos"},
		{12362, "cento e vinte e três reais e sessenta e dois centavos"},
		{5434021031, "cinquenta e quatro milhões, trezentos e quarenta mil, duzentos e dez reais e trinta e um centavos"},

		// The cases the original never had to answer.
		{0, ""},   // TODO: decide — "zero reais"?
		{100, ""}, // TODO: "um real"
		{-129, ""},
	}

	for _, tt := range tests {
		t.Run(tt.in.String(), func(t *testing.T) {
			if got := number.SpellMoney(tt.in); got != tt.want {
				t.Errorf("SpellMoney(%d) = %q, want %q", int64(tt.in), got, tt.want)
			}
		})
	}
}

func TestSpell(t *testing.T) {
	tests := []struct {
		in   int64
		want string
	}{
		{0, "zero"},
		{100, "cem"},        // not "cento"
		{101, "cento e um"},
		{1000, "mil"},       // not "um mil"
		{1200, ""},          // TODO: "mil e duzentos" or "mil duzentos"?
		{2000000, ""},       // TODO: "dois milhões"
		{-5, "menos cinco"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := number.Spell(tt.in); got != tt.want {
				t.Errorf("Spell(%d) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
