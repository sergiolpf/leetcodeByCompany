package longeststringchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestStrChain_TwoWordsChain(t *testing.T) {

	got := LongestStrChain([]string{"a", "ab"})
	want := 2

	assert.Equal(t, want, got)
}

func TestLongestStrChain_OneWordChain(t *testing.T) {

	got := LongestStrChain([]string{"a"})
	want := 1

	assert.Equal(t, want, got)
}

func TestLongestStrChain_NoWordChain(t *testing.T) {

	got := LongestStrChain([]string{"a", "b"})
	want := 1

	assert.Equal(t, want, got)
}

func TestLongestStrChain_MultipleWordsChain(t *testing.T) {
	listOfWords := []string{"a", "ab", "ac", "bd", "abc", "abd", "abdd"}
	got := LongestStrChain(listOfWords)
	want := 4

	assert.Equal(t, want, got)
}

func TestLongestStrChain_EmptyString(t *testing.T) {

	got := LongestStrChain([]string{""})
	want := 1

	assert.Equal(t, want, got)
}

func TestLongestStrChain_DuplicateWords(t *testing.T) {
	listOfWords := []string{"a", "ab", "ac", "bd", "abc", "abd", "abdd", "a", "ab", "ac"}
	got := LongestStrChain(listOfWords)
	want := 4

	assert.Equal(t, want, got)
}

func TestLongestStrChain_WordsWithDifferentCase(t *testing.T) {
	listOfWords := []string{"a", "ab", "AB", "aB", "abc"}
	got := LongestStrChain(listOfWords)
	want := 3

	assert.Equal(t, want, got)
}

func TestLongestStrChain_LongChainWithBranching(t *testing.T) {
	listOfWords := []string{"a", "ab", "abc", "abcd", "abd", "abde", "abcdef"}
	got := LongestStrChain(listOfWords)
	want := 4

	assert.Equal(t, want, got)
}
