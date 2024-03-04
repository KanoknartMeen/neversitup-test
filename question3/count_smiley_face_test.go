package question3

import (
	"testing"
)

func TestCountSmileys(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want int
	}{
		{
			name: "Test case 1",
			arr:  []string{":)", ";(", ";}", ":-D"},
			want: 2,
		},
		{
			name: "Test case 2",
			arr:  []string{";D", ":-(", ":-)", ";~)"},
			want: 3,
		},
		{
			name: "Test case 3",
			arr:  []string{";]", ":[", ";*", ":$", ";-D"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSmileys(tt.arr); got != tt.want {
				t.Errorf("CountSmileys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidSmiley(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{"Test case 1", ":)", true},
		{"Test case 2", ";D", true},
		{"Test case 3", ":-)", true},
		{"Test case 4", ";~)", true},
		{"Test case 5", ";-D", true},
		{"Test case 6", ":-D", true},
		{"Test case 7", ":D", true},
		{"Test case 8", ";(", false},
		{"Test case 9", ";}", false},
		{"Test case 10", ":-(", false},
		{"Test case 11", ";]", false},
		{"Test case 12", ":[", false},
		{"Test case 13", ";*", false},
		{"Test case 14", ":$", false},
		{"Test case 15", "", false},
		{"Test case 16", ":", false},
		{"Test case 17", ":-))", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSmiley(tt.str); got != tt.want {
				t.Errorf("isValidSmiley() = %v, want %v", got, tt.want)
			}
		})
	}
}
