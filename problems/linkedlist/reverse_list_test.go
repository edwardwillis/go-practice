package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty list", []int{}, nil},
		{"single node", []int{1}, []int{1}},
		{"multiple nodes", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := FromSlice(tt.in)
			got := ToSlice(Reverse(head))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
