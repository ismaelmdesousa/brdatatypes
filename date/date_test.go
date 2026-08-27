package date_test

import (
	"testing"
	"time"

	"github.com/ismaelmdesousa/brdatatypes/date"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		y       int
		m       time.Month
		d       int
		wantErr bool
	}{
		{name: "ordinary day", y: 2026, m: time.August, d: 24},
		{name: "last day of february, leap year", y: 2024, m: time.February, d: 29},

		// time.Date would silently return 2 March here. New must not.
		{name: "30 february", y: 2026, m: time.February, d: 30, wantErr: true},
		{name: "29 february, common year", y: 2026, m: time.February, d: 29, wantErr: true},
		{name: "31 april", y: 2026, m: time.April, d: 31, wantErr: true},
		{name: "day zero", y: 2026, m: time.August, d: 0, wantErr: true},
		{name: "month thirteen", y: 2026, m: time.Month(13), d: 1, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := date.New(tt.y, tt.m, tt.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("New(%d, %v, %d) error = %v, wantErr = %v", tt.y, tt.m, tt.d, err, tt.wantErr)
			}
		})
	}
}
