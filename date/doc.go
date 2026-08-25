// Package date provides a calendar date — a year, a month and a day, with no
// time of day and no timezone.
//
// It exists because time.Time carries an instant, and an instant is the wrong
// model for a birth date, a due date or a competence period: it drags timezone
// conversions into values that have none, and turns "2026-08-24" into a moment
// that can shift a day when it crosses a boundary.
package date
