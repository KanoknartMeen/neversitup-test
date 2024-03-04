package question2

import "testing"

func TestCountOddOccur(t *testing.T) {
	tests := []struct {
		name   string
		intArr []int
		want   int
	}{
		{
			name:   "Test case 1",
			intArr: []int{7},
			want:   7,
		},
		{
			name:   "Test case 2",
			intArr: []int{0},
			want:   0,
		},
		{
			name:   "Test case 3",
			intArr: []int{1, 1, 2},
			want:   2,
		},
		{
			name:   "Test case 4",
			intArr: []int{0, 1, 0, 1, 0},
			want:   0,
		},
		{
			name:   "Test case 5",
			intArr: []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			want:   4,
		},
		{
			name:   "Test case Fail",
			intArr: []int{2, 2},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountOddOccur(tt.intArr); got != tt.want {
				t.Errorf("CountOddOccur() = %v, want %v", got, tt.want)
			}
		})
	}
}
