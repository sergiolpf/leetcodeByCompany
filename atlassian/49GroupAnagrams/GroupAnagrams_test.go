package groupanagrams

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagramsWithOne_SingleWord(t *testing.T) {

	listOfWords := []string{
		"abc",
	}

	want := [][]string{
		{"abc"},
	}

	got := groupAnagrams(listOfWords)

	assert.ElementsMatch(t, want, got)
}
func TestGroupAnagrams_EmptyInput(t *testing.T) {

	listOfWords := []string{""}

	want := [][]string{
		{""},
	}

	got := groupAnagrams(listOfWords)

	assert.ElementsMatch(t, want, got)
}

func TestGroupAnagrams_SimpleAnagrams(t *testing.T) {

	listOfWords := []string{
		"eat", "tea", "ate",
	}

	want := [][]string{
		{"ate", "eat", "tea"},
	}

	got := groupAnagrams(listOfWords)

	// Ordena os anagramas dentro de cada grupo
	for _, group := range got {
		sort.Strings(group)
	}
	for _, group := range want {
		sort.Strings(group)
	}

	assert.ElementsMatch(t, want, got)
}

func TestGroupAnagrams_MultipleAnagrams(t *testing.T) {

	listOfWords := []string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}

	want := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}

	got := groupAnagrams(listOfWords)

	// Ordena os anagramas dentro de cada grupo
	for _, group := range got {
		sort.Strings(group)
	}
	for _, group := range want {
		sort.Strings(group)
	}

	assert.ElementsMatch(t, want, got)
}

func TestGroupAnagrams_NoAnagrams(t *testing.T) {

	listOfWords := []string{
		"hello", "world", "foo", "bar",
	}

	want := [][]string{
		{"hello"},
		{"world"},
		{"foo"},
		{"bar"},
	}

	got := groupAnagrams(listOfWords)

	// Ordena os anagramas dentro de cada grupo
	for _, group := range got {
		sort.Strings(group)
	}
	for _, group := range want {
		sort.Strings(group)
	}

	assert.ElementsMatch(t, want, got)
}
