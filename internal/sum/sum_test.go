package sum

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Positive numbers", 2, 3, 5},
		{"Negative numbers", -1, -5, -6},
		{"Mixed numbers", 5, -2, 3},
		{"Zero with positive", 0, 7, 7},
		{"Zero with negative", -4, 0, -4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Addition(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Addition(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Positive", 12, 6, 6},
	}

	for _, tt := range tests {
		got := Subtract(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Subtract(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
