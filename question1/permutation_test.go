package question1

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []string
	}{
		{
			name: "Permutation of 'a'",
			str:  "a",
			want: []string{"a"},
		},
		{
			name: "Permutation of 'ab'",
			str:  "ab",
			want: []string{"ab", "ba"},
		},
		{
			name: "Permutation of 'abc'",
			str:  "abc",
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name: "Permutation of 'aabb'",
			str:  "aabb",
			want: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
		{
			name: "Permutation of ''",
			str:  "",
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permutation(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicate(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want []string
	}{
		{
			name: "Test case 1",
			strs: []string{"a", "b", "c", "a", "b", "c"},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicate(tt.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
