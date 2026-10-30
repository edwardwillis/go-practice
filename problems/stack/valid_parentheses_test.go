package stack

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"simple valid", "()", true},
		{"mixed valid", "()[]{}", true},
		{"wrong order", "(]", false},
		{"unmatched open", "([)]", false},
		{"nested valid", "{[]}", true},
		{"empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.in); got != tt.want {
				t.Errorf("IsValid(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
