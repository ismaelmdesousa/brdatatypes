package taxid

import "database/sql/driver"

// CPF identifies an individual taxpayer: eleven digits, the last two of which
// are check digits.
//
// The value is held normalised — digits only, no mask — so two CPFs written
// differently compare equal. The zero value is the empty CPF; see IsZero.
type CPF struct {
	digits string
}

// ParseCPF reads a CPF from s, accepting it masked (000.000.000-00) or bare.
func ParseCPF(s string) (CPF, error) {
	panic("TODO: implement")
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

// String returns the masked form, 000.000.000-00, or the empty string for the
// zero value.
func (c CPF) String() string {
	panic("TODO: implement")
}

// Unmasked returns the eleven digits with no punctuation — the form to store in
// a database column.
func (c CPF) Unmasked() string { return c.digits }

// IsZero reports whether c is the zero value, meaning no CPF was informed.
func (c CPF) IsZero() bool { return c.digits == "" }

// MarshalText implements encoding.TextMarshaler. Because encoding/json falls
// back to TextMarshaler, this also decides how a CPF appears in JSON: bare
// digits, which is what an API consumer can store without stripping anything.
func (c CPF) MarshalText() ([]byte, error) {
	panic("TODO: implement")
}

// UnmarshalText implements encoding.TextUnmarshaler, accepting the same inputs
// as ParseCPF. Note that a JSON null leaves the value untouched, so a null
// decodes to the zero CPF.
func (c *CPF) UnmarshalText(text []byte) error {
	panic("TODO: implement")
}

// Value implements driver.Valuer, writing the bare digits, and NULL for the
// zero value.
func (c CPF) Value() (driver.Value, error) {
	panic("TODO: implement")
}

// Scan implements sql.Scanner, accepting string, []byte and nil, since drivers
// differ on which one they hand back for a text column.
func (c *CPF) Scan(src any) error {
	panic("TODO: implement")
}
