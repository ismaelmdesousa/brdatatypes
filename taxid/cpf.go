package taxid

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"slices"
)

var (
	invalidCPFs = []string{
		"00000000000",
		"11111111111",
		"22222222222",
		"33333333333",
		"44444444444",
		"55555555555",
		"66666666666",
		"77777777777",
		"88888888888",
		"99999999999",
	}
)

// CPF identifies an individual taxpayer: eleven digits, the last two of which
// are check digits.
//
// The value is held normalised — digits only, no mask — so two CPFs written
// differently compare equal. The zero value is the empty CPF; see IsZero.
type CPF struct {
	digits string
}

// calculateCPFCheckDigit calculates the check digit for a given CPF.
func calculateCPFCheckDigit(s string) (int, error) {
	weights := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0
	for i, char := range s {
		if char < '0' || char > '9' {
			return 0, ErrInvalidCharacter
		}
		sum += int(char-'0') * weights[len(weights)-len(s)+i]
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0, nil
	}
	return 11 - remainder, nil
}

// ParseCPF reads a CPF from s, accepting it masked (000.000.000-00) or bare.
func ParseCPF(s string) (CPF, error) {
	if ok, err := regexp.MatchString(`[^0-9\.\-]`, s); err != nil || ok {
		return CPF{}, ErrInvalidCharacter
	}

	cpf := regexp.MustCompile(`[^0-9]`).ReplaceAllString(s, "")
	if slices.Contains(invalidCPFs, cpf) {
		return CPF{}, ErrWellKnownInvalid
	}

	if len(cpf) != 11 {
		return CPF{}, ErrInvalidLength
	}

	dv1, err := calculateCPFCheckDigit(cpf[:9])
	if err != nil {
		return CPF{}, err
	} else if dv1 != int(cpf[9]-'0') {
		return CPF{}, ErrInvalidCheckDigit
	}

	dv2, err := calculateCPFCheckDigit(cpf[:10])
	if err != nil {
		return CPF{}, err
	} else if dv2 != int(cpf[10]-'0') {
		return CPF{}, ErrInvalidCheckDigit
	}

	return CPF{digits: cpf}, nil
}

// MustParseCPF is like ParseCPF but panics on invalid input. Reserve it for
// constants in tests and package-level values, never for user input.
func MustParseCPF(s string) CPF {
	c, err := ParseCPF(s)
	if err != nil {
		panic(err)
	}
	return c
}

// IsValidCPF reports whether s holds a valid CPF, without building a value.
func IsValidCPF(s string) bool {
	_, err := ParseCPF(s)
	return err == nil
}

// String returns the fourteen characters with no punctuation.
func (c CPF) String() string { return c.digits }

// Masked returns the masked form, 000.000.000-00, or the empty string for
// the zero value.
func (c CPF) Masked() string {
	return fmt.Sprintf("%s.%s.%s-%s", c.digits[0:3], c.digits[3:6], c.digits[6:9], c.digits[9:11])
}

// IsZero reports whether c is the zero value, meaning no CPF was informed.
func (c CPF) IsZero() bool { return c.digits == "" }

// MarshalText implements encoding.TextMarshaler. Because encoding/json falls
// back to TextMarshaler, this also decides how a CPF appears in JSON: bare
// digits, which is what an API consumer can store without stripping anything.
func (c CPF) MarshalText() ([]byte, error) {
	return []byte(c.digits), nil
}

// UnmarshalText implements encoding.TextUnmarshaler, accepting the same inputs
// as ParseCPF. Note that a JSON null leaves the value untouched, so a null
// decodes to the zero CPF.
func (c *CPF) UnmarshalText(text []byte) error {
	cpf, err := ParseCPF(string(text))
	if err != nil {
		return err
	}
	*c = cpf
	return nil
}

// Value implements driver.Valuer, writing the bare digits, and NULL for the
// zero value.
func (c CPF) Value() (driver.Value, error) {
	if c.IsZero() {
		return nil, nil
	}
	return c.digits, nil
}

// Scan implements sql.Scanner, accepting string, []byte and nil, since drivers
// differ on which one they hand back for a text column.
func (c *CPF) Scan(src any) error {
	switch v := src.(type) {
	case nil:
		*c = CPF{}
		return nil
	case string:
		cpf, err := ParseCPF(v)
		if err != nil {
			return err
		}
		*c = cpf
		return nil
	case []byte:
		cpf, err := ParseCPF(string(v))
		if err != nil {
			return err
		}
		*c = cpf
		return nil
	default:
		return fmt.Errorf("cannot scan %T into CPF", src)
	}
}
