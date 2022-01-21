package search_test

import (
	"testing"

	"github.com/yousifsabah0/Golorithms/search"
)

func TestLinear(t *testing.T) {
	var tests = []struct {
		arr          []int
		target, want int
	}{
		{[]int{0, 20, 80, 30, 60, 50, 110, 100, 130, 170}, 110, 6},
		{[]int{0, 20, 80, 30, 60, 50, 110, 100, 130, 170}, 200, -1},
	}
	for _, test := range tests {
		got := search.Linear(test.arr, test.target)
		if got != test.want {
			t.Errorf("got %q; want %q", got, test.want)
		}
	}
}
