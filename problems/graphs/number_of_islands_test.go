package graphs

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "single island",
			grid: [][]byte{
				[]byte("11110"),
				[]byte("11010"),
				[]byte("11000"),
				[]byte("00000"),
			},
			want: 1,
		},
		{
			name: "three islands",
			grid: [][]byte{
				[]byte("11000"),
				[]byte("11000"),
				[]byte("00100"),
				[]byte("00011"),
			},
			want: 3,
		},
		{
			name: "empty grid",
			grid: [][]byte{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumIslands(tt.grid); got != tt.want {
				t.Errorf("NumIslands() = %d, want %d", got, tt.want)
			}
		})
	}
}
