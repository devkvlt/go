package runes

import "testing"

type test struct {
	name  string
	input rune
	want  bool
}

func TestIsHex(t *testing.T) {
	tests := []test{
		{"Test 1", 'a', true},
		{"Test 2", 'A', true},
		{"Test 3", '4', true},
		{"Test 4", 'y', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			got := IsHex(tt.input)
			if got != want {
				t.Errorf("\n\twant\t%v\n\tgot\t%v", want, got)
			}
		})
	}
}
