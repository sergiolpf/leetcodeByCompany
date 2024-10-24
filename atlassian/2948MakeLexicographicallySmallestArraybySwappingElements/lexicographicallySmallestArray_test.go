package lexicographicallySmallestArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexicographicallySmallestArray_SimpleSwaps(t *testing.T) {
	got := lexicographicallySmallestArray([]int{1, 5, 3, 9, 8}, 2)

	want := []int{1, 3, 5, 8, 9}

	assert.ElementsMatch(t, want, got)
}
func TestLexicographicallySmallestArray_NoSwaps(t *testing.T) {
	got := lexicographicallySmallestArray([]int{1, 7, 6, 18, 2, 1}, 3)

	want := []int{1, 6, 7, 18, 1, 2}

	assert.ElementsMatch(t, want, got)
}

func TestLexicographicallySmallestArray_EmptyArray(t *testing.T) {
	got := lexicographicallySmallestArray([]int{}, 2)
	want := []int{}
	assert.ElementsMatch(t, want, got)
}

func TestLexicographicallySmallestArray_SingleElement(t *testing.T) {
	got := lexicographicallySmallestArray([]int{5}, 2)
	want := []int{5}
	assert.ElementsMatch(t, want, got)
}

func TestLexicographicallySmallestArray_AllElementsSameGroup(t *testing.T) {
	got := lexicographicallySmallestArray([]int{2, 3, 1, 4}, 1)
	want := []int{1, 2, 3, 4}
	assert.ElementsMatch(t, want, got)
}

func TestLexicographicallySmallestArray_LargeLimit(t *testing.T) {
	got := lexicographicallySmallestArray([]int{5, 1, 9, 3}, 10)
	want := []int{1, 3, 5, 9}
	assert.ElementsMatch(t, want, got)
}
