package relativetime_test

import (
	"testing"
	"time"

	"github.com/abdelkd/todo-cli/internal/relativetime"
)

func TestRelativePackage(t *testing.T) {
	now := time.Now()
	tests := []struct {
		pastTime       time.Time
		expectedString string
	}{
		{now.Add(-10 * time.Second), "just now"},

		{now.Add(-1 * time.Minute), "1 minute ago"},
		{now.Add(-3 * time.Minute), "3 minutes ago"},

		{now.Add(-1 * time.Hour), "1 hour ago"},
		{now.Add(-4 * time.Hour), "4 hours ago"},

		{now.Add(-25 * time.Hour), "1 day ago"},
		{now.Add(-144 * time.Hour), "6 days ago"},

		{now.Add(-169 * time.Hour), "1 week ago"},
		{now.Add(-507 * time.Hour), "3 weeks ago"},

		{now.Add(-676 * time.Hour), "1 month ago"},
		{now.Add(-2028 * time.Hour), "3 months ago"},

		{now.Add(-676 * time.Hour), "1 year ago"},
		{now.Add(-2028 * time.Hour), "3 years ago"},
	}

	for _, test := range tests {
		result := relativetime.RelativeTime(test.pastTime)
		if result != test.expectedString {
			// now := time.Now()
			// diff := now.Sub(test.pastTime)
			// panic(diff >= (time.Hour*169) && diff <= ((time.Hour*169)*4))
			t.Fatalf("incorrect time string, expected=%s got=%s", result, test.expectedString)
		}
	}
}
