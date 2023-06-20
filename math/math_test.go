package math

import "testing"

type test struct {
	name  string
	input []int
	want  int
}

func TestMinInt(t *testing.T) {
	tests := []test{
		{"Test 1", []int{1, 2, 3}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			got := Min(tt.input...)
			if got != want {
				t.Errorf("\n\twant\t%v\n\tgot\t%v", want, got)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	tests := []test{
		{"Test 1", []int{1, 2, 3}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			got := Max(tt.input...)
			if got != want {
				t.Errorf("\n\twant\t%v\n\tgot\t%v", want, got)
			}
		})
	}
}

func TestSumInt(t *testing.T) {
	tests := []test{
		{"Test 1", []int{1, 2, 3}, 6},
		{"Test 2", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			got := Sum(tt.input...)
			if got != want {
				t.Errorf("\n\twant\t%v\n\tgot\t%v", want, got)
			}
		})
	}
}
