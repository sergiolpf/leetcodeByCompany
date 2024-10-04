package maximumareaoflongestdiagonalrectangle

import "testing"

func TestAreaOfMaxDiagonal(t *testing.T) {
	got := AreaOfMaxDiagonal([][]int{
		{6, 5}, {8, 6}, {2, 10}, {8, 1}, {9, 2}, {3, 5}, {3, 5},
	})
	expect := 12

	if got != expect {
		t.Fail()
	}
}
