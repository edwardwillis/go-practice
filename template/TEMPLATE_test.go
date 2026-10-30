package template

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{"case 1", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.in); got != tt.want {
				t.Errorf("Solve(%v) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
