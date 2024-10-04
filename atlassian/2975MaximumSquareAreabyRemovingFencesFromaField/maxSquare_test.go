package maximumsquareareabyremovingfencesfromafield

import (
	"testing"
)

func TestMaximizeSquareArea(t *testing.T) {

	// got := MaxSquareField(4, 3, []int{2, 3}, []int{2})
	got := MaximizeSquareArea(6, 7, []int{2}, []int{4})
	want := 4

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
