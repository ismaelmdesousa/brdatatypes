package taxid

// Kind tells which identifier a value holds.
type Kind int

const (
	// Unknown is the zero Kind: the input matched neither identifier.
	Unknown Kind = iota
	// KindCPF is an individual taxpayer identifier.
	KindCPF
	// KindCNPJ is a company taxpayer identifier.
	KindCNPJ
)

// Parse reads whichever identifier s holds, deciding by the number of
// significant characters: eleven means CPF, fourteen means CNPJ.
//
// Use it where the source field accepts both — an ERP column that stores the
// taxpayer of a customer who may be a person or a company. Where the domain
// admits only one, prefer ParseCPF or ParseCNPJ so the compiler keeps them
// apart.
func Parse(s string) (Kind, CPF, CNPJ, error) {
	panic("TODO: implement")
}
