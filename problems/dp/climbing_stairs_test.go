package dp

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}

	for _, tt := range tests {
		if got := ClimbStairs(tt.n); got != tt.want {
			t.Errorf("ClimbStairs(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
