package taxid

import "errors"

// Errors returned by the parsers. Compare with errors.Is rather than by string.
var (
	// ErrInvalidLength means the input does not have the number of significant
	// characters the identifier requires: eleven for CPF, fourteen for CNPJ.
	ErrInvalidLength = errors.New("taxid: invalid length")

	// ErrInvalidCharacter means the input carries a character that is neither
	// punctuation nor valid in that position.
	ErrInvalidCharacter = errors.New("taxid: invalid character")

	// ErrInvalidCheckDigit means the input is well formed but the check digits do
	// not match the rest of the value.
	ErrInvalidCheckDigit = errors.New("taxid: invalid check digit")

	// ErrWellKnownInvalid means the input passes the check digit arithmetic but is
	// rejected by convention, such as a CPF whose eleven digits are all the same.
	ErrWellKnownInvalid = errors.New("taxid: well-known invalid value")
)
