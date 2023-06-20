package slices

import (
	"reflect"
	"testing"
)

func TestEqual(t *testing.T) {
	type test[T comparable] struct {
		name string
		s1   []T
		s2   []T
		want bool
	}

	test1 := test[int]{
		name: "test int",
		s1:   []int{1, 2, 3},
		s2:   []int{1, 2, 3},
		want: true,
	}

	test2 := test[string]{
		name: "test string",
		s1:   []string{"1", "2", "3"},
		s2:   []string{"1", "2", "3"},
		want: true,
	}

	t.Run(test1.name, func(t *testing.T) {
		if got := Equal(test1.s1, test1.s2); got != test1.want {
			t.Errorf("Equal() = %v, want %v", got, test1.want)
		}
	})

	t.Run(test2.name, func(t *testing.T) {
		if got := Equal(test2.s1, test2.s2); got != test2.want {
			t.Errorf("Equal() = %v, want %v", got, test2.want)
		}
	})
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "Test 1",
			s:    []int{1, 2, 2, 3, 4, 3, 4, 1},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Test 2",
			s:    []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
