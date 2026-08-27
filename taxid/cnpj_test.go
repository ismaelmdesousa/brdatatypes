package taxid_test

import (
	"errors"
	"testing"

	"github.com/ismaelmdesousa/brdatatypes/taxid"
)

func TestParseCNPJ(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr error
	}{
		// numeric values, taken from real published identifiers.
		{name: "numeric masked", in: "48.071.019/0001-55", want: "48071019000155"},

		// alphanumeric values, generated in the Receita Federal simulator
		// and pinned here. Do not invent these — the check digits must come from
		// the official implementation.
		{name: "alphanumeric", in: "HY.DK2.BWL/0001-95", want: "HYDK2BWL000195"},
		{name: "lowercase is accepted and uppercased", in: "hy.dk2.bwl/0001-95", want: "HYDK2BWL000195"},

		// The last two characters are check digits and are always numeric.
		{name: "letter in check digit", in: "12ABC34501DEAB", wantErr: taxid.ErrInvalidCheckDigit},

		{name: "too short", in: "1234567890123", wantErr: taxid.ErrInvalidLength},
		{name: "empty", in: "", wantErr: taxid.ErrInvalidLength},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := taxid.ParseCNPJ(tt.in)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("ParseCNPJ(%q) error = %v, want %v", tt.in, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseCNPJ(%q) unexpected error: %v", tt.in, err)
			}
			if got.String() != tt.want {
				t.Errorf("ParseCNPJ(%q) = %q, want %q", tt.in, got.String(), tt.want)
			}
		})
	}
}
