package taxid_test

import (
	"errors"
	"testing"

	"github.com/ismaelmdesousa/brdatatypes/taxid"
)

func TestParseCPF(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string // unmasked; empty when an error is expected
		wantErr error
	}{
		// TODO: fill in with valid values. Generate them rather than inventing
		// them by hand — a hand-written CPF almost never has correct check digits.
		{name: "masked", in: "", want: ""},
		{name: "bare", in: "", want: ""},

		// The eleven repeated-digit values pass modulus 11 arithmetic and must be
		// rejected by rule. This is the case a naive implementation gets wrong.
		{name: "all ones", in: "11111111111", wantErr: taxid.ErrWellKnownInvalid},
		{name: "all zeros", in: "00000000000", wantErr: taxid.ErrWellKnownInvalid},

		{name: "too short", in: "1234567890", wantErr: taxid.ErrInvalidLength},
		{name: "letter", in: "1234567890A", wantErr: taxid.ErrInvalidCharacter},
		{name: "empty", in: "", wantErr: taxid.ErrInvalidLength},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := taxid.ParseCPF(tt.in)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("ParseCPF(%q) error = %v, want %v", tt.in, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseCPF(%q) unexpected error: %v", tt.in, err)
			}
			if got.Unmasked() != tt.want {
				t.Errorf("ParseCPF(%q) = %q, want %q", tt.in, got.Unmasked(), tt.want)
			}
		})
	}
}

// FuzzParseCPF checks the round trip: anything ParseCPF accepts must survive
// being formatted and parsed again unchanged.
func FuzzParseCPF(f *testing.F) {
	f.Add("529.982.247-25")
	f.Add("11111111111")
	f.Fuzz(func(t *testing.T, s string) {
		c, err := taxid.ParseCPF(s)
		if err != nil {
			return
		}
		again, err := taxid.ParseCPF(c.String())
		if err != nil {
			t.Fatalf("ParseCPF(%q) accepted, but its String() %q was rejected: %v", s, c.String(), err)
		}
		if again != c {
			t.Errorf("round trip changed the value: %q -> %q", c.Unmasked(), again.Unmasked())
		}
	})
}
