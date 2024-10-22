package smallestmissing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSmallestInteger(t *testing.T) {
	got := findSmallestInteger([]int{1, -10, 7, 13, 6, 8}, 5)
	want := 4

	assert.Equal(t, want, got)

	got = findSmallestInteger([]int{1, -10, 7, 13, 6, 8}, 7)
	want = 2

	assert.Equal(t, want, got)
}
package smallestmissing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSmallestInteger_EmptyArray(t *testing.T) {
	got := findSmallestInteger([]int{}, 5)
	want := 0

	assert.Equal(t, want, got)
}

func TestFindSmallestInteger_AllPositive(t *testing.T) {
	got := findSmallestInteger([]int{1, 2, 3, 4}, 5)
	want := 5

	assert.Equal(t, want, got)
}

func TestFindSmallestInteger_AllNegative(t *testing.T) {
	got := findSmallestInteger([]int{-1, -2, -3, -4}, 5)
	want := 1

	assert.Equal(t, want, got)
}

func TestFindSmallestInteger_PositiveAndNegative(t *testing.T) {
	ot := findSmallestInteger([]int{1, -10, 7, 13, 6, 8}, 5)
	want := 4

	got := findSmallestInteger([]int{-1, 2, -3, 4}, 5)
	want := 3

	assert.Equal(t, want, got)
}

func TestFindSmallestInteger_LargeValue(t *testing.T) {
	got := findSmallestInteger([]int{1, 2, 3, 4}, 10)
	want := 5

	assert.Equal(t, want, got)
}

func TestFindSmallestInteger_DuplicateNumbers(t *testing.T) {
	got := findSmallestInteger([]int{1, 2, 2, 2, 3, 3, 4, 4, 4, 4}, 5)
	want := 5

	assert.Equal(t, want, got)
}



