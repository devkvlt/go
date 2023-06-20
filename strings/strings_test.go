package strings

import "testing"

type test struct {
	name  string
	input string
	want  bool
}

func TestIsHex(t *testing.T) {
	tests := []test{
		{"Test 1", "abc", true},
		{"Test 2", "AAAA", true},
		{"Test 3", "4aF", true},
		{"Test 4", "Yo", false},
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
