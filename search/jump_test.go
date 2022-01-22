package search_test

import (
	"testing"

	"github.com/yousifsabah0/Golorithms/search"
)

func TestJump(t *testing.T) {
	var tests = []struct {
		arr          []int
		target, want int
	}{
		{[]int{0, 20, 30, 50, 60, 70, 80, 100, 130, 170}, 110, 6},
		{[]int{0, 20, 30, 50, 60, 70, 80, 100, 130, 170}, 200, -1},
	}
	for _, test := range tests {
		got := search.Jump(test.arr, test.target)
		if got != test.want {
			t.Errorf("got %q; want %q", got, test.want)
		}
	}
}
