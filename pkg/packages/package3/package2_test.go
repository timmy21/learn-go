package package3

import (
	"testing"
)

func TestNums_Sum(t *testing.T) {
	tests := []struct {
		name string
		n    Nums
		want int
	}{
		{
			name: "basic",
			n:    []int{1, 2, 3},
			want: 6,
		},
		{
			name: "empty",
			n:    nil,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Sum(); got != tt.want {
				t.Errorf("Nums.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNums_Avg(t *testing.T) {
	tests := []struct {
		name string
		n    Nums
		want int
	}{
		{
			name: "basic",
			n:    []int{1, 2, 3},
			want: 2,
		},
		{
			name: "empty",
			n:    nil,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Avg(); got != tt.want {
				t.Errorf("Nums.Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}
