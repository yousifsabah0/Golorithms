package problems_test

import (
	"testing"

	"github.com/yousifsabah0/Golorithms/problems"
)

func TestFindHighestElement(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{0, 20, 30, 50, 60, 70, 80, 100, 130, 170}, 170},
	}
	for _, test := range tests {
		got := problems.FindHighestElement(test.arr)
		if got != test.want {
			t.Errorf("got %q; want %q", got, test.want)
		}
	}
}

func TestFindSecondHighestElement(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{0, 20, 30, 50, 60, 70, 80, 100, 130, 170}, 130},
	}
	for _, test := range tests {
		got := problems.FindSecondHighestElement(test.arr)
		if got != test.want {
			t.Errorf("got %q; want %q", got, test.want)
		}
	}
}
