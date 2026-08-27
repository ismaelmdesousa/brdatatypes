# Design decisions

Decisions that shape the public API, written down so they stop being re-argued.
Each records what was chosen and what it costs, not only what it gains.

All five are **settled**. Code may rely on them.

---

## 1. A single type, not one per identifier

`taxid.TaxID` holds either a CPF or a CNPJ. There is no separate `CPF` type and
no separate `CNPJ` type.

```go
type TaxID struct {
	value string // normalised: no mask, letters uppercased
}
```

The kind is **derived, never stored** — eleven significant characters is a CPF,
fourteen is a CNPJ. A stored `kind` field could disagree with the value it
describes; deriving it makes that state unrepresentable, and keeps `TaxID`
comparable with `==` and usable as a map key.

**Why.** The column this maps to usually accepts both, because a customer may be
a person or a company. One type lets that column round trip without the caller
branching first, and keeps a single implementation of the JSON and SQL plumbing.

**What it costs.** The compiler no longer stops a CPF reaching a parameter that
expects a CNPJ. That check moves to runtime, so the constructors carry it:
`ParseCPF` and `ParseCNPJ` parse *and* assert the kind, rejecting the other.

```go
func Parse(s string) (TaxID, error)      // accepts either, decides by length
func ParseCPF(s string) (TaxID, error)   // rejects a valid CNPJ
func ParseCNPJ(s string) (TaxID, error)  // rejects a valid CPF

func (t TaxID) Kind() Kind    // KindCPF | KindCNPJ
func (t TaxID) IsCPF() bool
func (t TaxID) IsCNPJ() bool
func (t TaxID) String() string   // mask matching the kind
func (t TaxID) Unmasked() string
```

---

## 2. The value is held normalised

Significant characters only — no punctuation, letters uppercased. The mask is
applied in `String()`, never stored.

Two identifiers written differently therefore compare equal, so `==` and map
keys behave the way a reader expects.

The cost is that the original formatting is lost. A caller who must echo input
back exactly as it arrived keeps that string themselves; this package answers
what the identifier *is*, not how it was typed.

---

## 3. JSON is asymmetric on purpose

**Reading** accepts masked and bare, because both arrive in real payloads.
**Writing** emits bare, because that is what a consumer can store without
stripping anything.

Implementing `MarshalText` / `UnmarshalText` is enough: `encoding/json` falls
back to them, so `MarshalJSON` / `UnmarshalJSON` would be duplicate surface.

---

## 4. The zero value is invalid, not "absent"

`TaxID{}` is not a taxpayer identifier. It is what you get from a declaration
that was never parsed, and the package treats it as an error rather than as a
blank that happens to be acceptable.

This holds for every type in the module — `date.Date{}` is likewise not a date.
`number.Money` is the exception, and not really one: zero cents is a legitimate
amount, not a missing value.

**Why.** `Parse` is the only way to build a valid value, so every value that
exists has been validated. Letting the zero value mean "fine, just empty" would
reintroduce the unvalidated state that parsing exists to remove, and it would
travel silently — an unset field would serialise, persist and compare as if it
had been informed.

**Consequences, which are the point of writing this down:**

| Situation | Behaviour |
| --- | --- |
| `TaxID{}.String()` | empty string — it has nothing to render |
| `TaxID{}.IsZero()` | `true`, meaning "never parsed", not "informed as blank" |
| `TaxID{}.Value()` | **error**. Writing an unvalidated value to a column is a bug, and failing loudly beats persisting a silent blank |
| `Scan(nil)` | **error**. A `NULL` cannot become a valid `TaxID` |
| `UnmarshalText("")` | error, same reason |
| JSON `null` | never reaches `UnmarshalText`; the field stays the zero value, which is invalid |

**How to express "not informed", then.** Use a pointer: `*TaxID`, `nil` when
absent. It is the idiom the standard library already uses for a nullable value
whose zero is meaningful, it costs nothing at the call site, and it makes
absence visible in the type instead of hidden inside it.

```go
type Customer struct {
	Name  string
	Tax   *taxid.TaxID `json:"tax_id"` // nil when the customer has not given one
}
```

---

## 5. Money is an integer, never a float

Currency is `number.Money`, an `int64` count of cents. Formatting and spelling
take `Money`; neither accepts `float64`.

`0.1 + 0.2` is not `0.3` in binary floating point. In a formatted figure that is
a rounding artefact someone might not notice. In an amount spelled out in full
words it becomes wrong text inside a contract — which is exactly what the
spelling function is used for.

`float64` remains available through `FormatFloat` and `ParseFloat`, for
measurements and quantities where binary rounding is acceptable.

---

## Conventions

- **No dependencies.** The standard library covers everything here, and "no
  dependencies" is itself a reason to adopt a library this small.
- **Sentinel errors**, compared with `errors.Is`, never by string matching.
- **Table-driven tests**, with fuzzing on the parsers — untrusted input is the
  entire job of a parser.
- **Go 1.22** in `go.mod`. A library should not force its consumers to upgrade.
