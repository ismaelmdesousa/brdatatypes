package taxid

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"strings"
)

// CNPJ identifies a company taxpayer: fourteen characters, the last two of
// which are numeric check digits.
//
// Since 6 July 2026 the first twelve characters may be digits or uppercase
// letters. The check digits use the same modulus 11 arithmetic as before, with
// each character worth ASCII(c) - 48: digits 0-9 keep their natural value and
// letters A-Z are worth 17 to 42. Identifiers issued before that date are
// unaffected and validate under the same rule.
//
// The value is held normalised — no mask, letters uppercased. The zero value is
// the empty CNPJ; see IsZero.
type CNPJ struct {
	chars string
}

// calculateCNPJCheckDigit computes the check digit for the first twelve characters of a CNPJ.
func calculateCNPJCheckDigit(chars string) (int, error) {
	weights := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum := 0
	for i, c := range chars {
		sum += int(c-'0') * weights[12-len(chars)+i+1]
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0, nil
	}
	return 11 - remainder, nil
}

// ParseCNPJ reads a CNPJ from s, accepting it masked (00.000.000/0000-00) or
// bare, with letters in either case.
func ParseCNPJ(s string) (CNPJ, error) {
	rx := regexp.MustCompile(`[^0-9A-Za-z]`)
	cnpj := strings.ToUpper(rx.ReplaceAllString(s, ""))

	if len(cnpj) != 14 {
		return CNPJ{}, ErrInvalidLength
	}

	for _, c := range cnpj[:12] {
		if !((c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z')) {
			return CNPJ{}, ErrInvalidCharacter
		}
	}

	if cnpj[12] < '0' || cnpj[12] > '9' || cnpj[13] < '0' || cnpj[13] > '9' {
		return CNPJ{}, ErrInvalidCheckDigit
	}

	dv1, err := calculateCNPJCheckDigit(cnpj[:12])
	if err != nil {
		return CNPJ{}, err
	} else if dv1 != int(cnpj[12]-'0') {
		return CNPJ{}, ErrInvalidCheckDigit
	}

	dv2, err := calculateCNPJCheckDigit(cnpj[:13])
	if err != nil {
		return CNPJ{}, err
	} else if dv2 != int(cnpj[13]-'0') {
		return CNPJ{}, ErrInvalidCheckDigit
	}

	return CNPJ{chars: cnpj}, nil
}

// MustParseCNPJ is like ParseCNPJ but panics on invalid input.
func MustParseCNPJ(s string) CNPJ {
	c, err := ParseCNPJ(s)
	if err != nil {
		panic(err)
	}
	return c
}

// IsValidCNPJ reports whether s holds a valid CNPJ, without building a value.
func IsValidCNPJ(s string) bool {
	_, err := ParseCNPJ(s)
	return err == nil
}

// IsAlphanumeric reports whether c uses any letter, which tells apart an
// identifier issued under the 2026 rule from a purely numeric one.
func (c CNPJ) IsAlphanumeric() bool {
	rx := regexp.MustCompile(`[a-zA-Z]`)
	return rx.MatchString(c.chars)
}

// Root returns the first eight characters, shared by every branch of the same
// company.
func (c CNPJ) Root() string {
	return c.chars[:8]
}

// String returns the fourteen characters with no punctuation.
func (c CNPJ) String() string { return c.chars }

// Masked returns the masked form, 00.000.000/0000-00, or the empty string for
// the zero value.
func (c CNPJ) Masked() string {
	return fmt.Sprintf("%s.%s.%s/%s-%s", c.chars[0:2], c.chars[2:5], c.chars[5:8], c.chars[8:12], c.chars[12:14])
}

// IsZero reports whether c is the zero value, meaning no CNPJ was informed.
func (c CNPJ) IsZero() bool { return c.chars == "" }

// MarshalText implements encoding.TextMarshaler, and by extension decides the
// JSON form: bare characters, no mask.
func (c CNPJ) MarshalText() ([]byte, error) {
	return []byte(c.Masked()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler, accepting the same inputs
// as ParseCNPJ.
func (c *CNPJ) UnmarshalText(text []byte) error {
	cnpj, err := ParseCNPJ(string(text))
	if err != nil {
		return err
	}
	*c = cnpj
	return nil
}

// Value implements driver.Valuer, writing the bare characters, and NULL for the
// zero value.
func (c CNPJ) Value() (driver.Value, error) {
	if c.IsZero() {
		return nil, nil
	}
	return c.String(), nil
}

// Scan implements sql.Scanner, accepting string, []byte and nil.
func (c *CNPJ) Scan(src any) error {
	switch v := src.(type) {
	case nil:
		*c = CNPJ{}
		return nil
	case string:
		cnpj, err := ParseCNPJ(v)
		if err != nil {
			return err
		}
		*c = cnpj
		return nil
	case []byte:
		cnpj, err := ParseCNPJ(string(v))
		if err != nil {
			return err
		}
		*c = cnpj
		return nil
	default:
		return fmt.Errorf("cannot scan %T into CNPJ", src)
	}
}
