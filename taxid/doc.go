// Package taxid handles Brazilian taxpayer identifiers: CPF for individuals and
// CNPJ for companies.
//
// Both are parsed from any input formatting, held normalised, and rendered back
// with the conventional mask on demand. The types implement the JSON and SQL
// interfaces so they can cross an API boundary and a database column without
// conversion code at every call site.
//
// CNPJ support covers the alphanumeric format that became valid on 6 July 2026:
// the first twelve characters may be digits or uppercase letters, the last two
// remain numeric check digits.
package taxid
