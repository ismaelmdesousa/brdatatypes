package date

import (
	"database/sql/driver"
	"time"
)

// Date is a calendar date: a year, a month and a day, with no time of day and
// no timezone.
//
// The zero value is the zero date, reported by IsZero, and is not a valid
// calendar date.
type Date struct {
	year  int
	month time.Month
	day   int
}

// New builds a Date, rejecting a day that does not exist in that month.
//
// Note this cannot be delegated to time.Date, which normalises silently:
// time.Date(2026, 2, 30, ...) returns 2 March rather than an error.
func New(year int, month time.Month, day int) (Date, error) {
	panic("TODO: implement")
}

// Parse reads a date from s, accepting ISO 2006-01-02 and 02/01/2006.
func Parse(s string) (Date, error) {
	panic("TODO: implement")
}

// FromTime discards the time of day of t, keeping the calendar date as it reads
// in t's own location.
func FromTime(t time.Time) Date {
	panic("TODO: implement")
}

// Today returns the current date in loc. A location is required because "today"
// is a different day either side of a timezone.
func Today(loc *time.Location) Date {
	return FromTime(time.Now().In(loc))
}

// Year, Month and Day return the components.
func (d Date) Year() int        { return d.year }
func (d Date) Month() time.Month { return d.month }
func (d Date) Day() int          { return d.day }

// Time places the date at midnight in loc.
func (d Date) Time(loc *time.Location) time.Time {
	panic("TODO: implement")
}

// String returns the ISO form, 2006-01-02.
func (d Date) String() string {
	panic("TODO: implement")
}

// Format returns the date in the Brazilian form, 02/01/2006.
func (d Date) Format() string {
	panic("TODO: implement")
}

// IsZero reports whether d is the zero value.
func (d Date) IsZero() bool { return d.year == 0 && d.month == 0 && d.day == 0 }

// Compare returns a negative number if d is before o, zero if they are the same
// date, and a positive number if d is after o.
func (d Date) Compare(o Date) int {
	panic("TODO: implement")
}

// Before and After compare two dates.
func (d Date) Before(o Date) bool { return d.Compare(o) < 0 }
func (d Date) After(o Date) bool  { return d.Compare(o) > 0 }

// AddDays returns the date n days away, where n may be negative.
func (d Date) AddDays(n int) Date {
	panic("TODO: implement")
}

// AddMonths returns the date n months away, clamping the day to the last day of
// the target month: 31 January plus one month is 28 or 29 February, never
// 2 or 3 March.
func (d Date) AddMonths(n int) Date {
	panic("TODO: implement")
}

// Sub returns the number of days between d and o.
func (d Date) Sub(o Date) int {
	panic("TODO: implement")
}

// MarshalText implements encoding.TextMarshaler, which also decides the JSON
// form: the ISO date, "2026-08-24".
func (d Date) MarshalText() ([]byte, error) {
	panic("TODO: implement")
}

// UnmarshalText implements encoding.TextUnmarshaler, accepting the same inputs
// as Parse.
func (d *Date) UnmarshalText(text []byte) error {
	panic("TODO: implement")
}

// Value implements driver.Valuer, writing a time.Time at midnight UTC, and NULL
// for the zero value.
func (d Date) Value() (driver.Value, error) {
	panic("TODO: implement")
}

// Scan implements sql.Scanner, accepting time.Time, string, []byte and nil,
// since drivers disagree on what a DATE column comes back as.
func (d *Date) Scan(src any) error {
	panic("TODO: implement")
}
