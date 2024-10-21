package alloonedatastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInc_AddingFirstElement(t *testing.T) {
	builtAllOne := Constructor()

	builtAllOne.Inc("aa")

	got := builtAllOne.GetMaxKey()
	want := "aa"

	assert.Equal(t, want, got)

	got = builtAllOne.GetMinKey()
	want = "aa"
	assert.Equal(t, want, got)
}

func TestInc_AddingMultipleRepeatedKeys(t *testing.T) {
	builtAllOne := Constructor()

	builtAllOne.Inc("dd")
	builtAllOne.Inc("dd")
	builtAllOne.Inc("dd")

	got := builtAllOne.GetMaxKey()
	want := "dd"

	assert.Contains(t, want, got)

	got = builtAllOne.GetMinKey()
	want = "dd"
	assert.Equal(t, want, got)
}

func TestInc_AddingMultipleUniqueAndRepeatedKeys(t *testing.T) {
	builtAllOne := Constructor()

	builtAllOne.Inc("aa")
	builtAllOne.Inc("bb")
	builtAllOne.Inc("cc")
	builtAllOne.Inc("dd")
	builtAllOne.Inc("dd")

	got := builtAllOne.GetMaxKey()
	want := "dd"

	assert.Equal(t, want, got)

	wantArray := []string{"aa", "bb", "cc"}
	got = builtAllOne.GetMinKey()
	assert.Contains(t, wantArray, got)

}

func TestGetMaxKey_ZeroInput(t *testing.T) {
	builtAllOne := Constructor()

	got := builtAllOne.GetMaxKey()
	want := ""

	assert.Equal(t, want, got)

}
func TestGetMinKey_ZeroInput(t *testing.T) {
	builtAllOne := Constructor()

	got := builtAllOne.GetMinKey()
	want := ""

	assert.Equal(t, want, got)

}

func TestDec_RemovingOneStringWithFrequenceOne(t *testing.T) {
	builtAllOne := Constructor()
	builtAllOne.Inc("aa")
	builtAllOne.Dec("aa")

	got := builtAllOne.GetMaxKey()
	want := ""
	assert.Equal(t, want, got)

	gotLen := len(builtAllOne.stringMap)
	wantLen := 0
	assert.Equal(t, wantLen, gotLen)

}
func TestDec_RemovingOneStringWithFrequenceOneWhenThereAreOtherStringsWithFrequenceOne(t *testing.T) {
	builtAllOne := Constructor()
	builtAllOne.Inc("aa")
	builtAllOne.Inc("Will")

	builtAllOne.Dec("aa")

	got := builtAllOne.GetMaxKey()
	want := "Will"
	assert.Equal(t, want, got)

	gotLen := len(builtAllOne.stringMap)
	wantLen := 1
	assert.Equal(t, wantLen, gotLen)

	_, ok := builtAllOne.stringMap["aa"]

	assert.False(t, ok)

}

func TestDec_RemovingOneStringWithFrequenceTwoWhenThereAreOtherStringsWithFrequenceOne(t *testing.T) {
	builtAllOne := Constructor()
	builtAllOne.Inc("aa")
	builtAllOne.Inc("aa")
	builtAllOne.Inc("Will")

	got := builtAllOne.GetMaxKey()
	want := "aa"
	assert.Equal(t, want, got)

	builtAllOne.Dec("aa")
	got = builtAllOne.GetMaxKey()
	wantArray := []string{"Will", "aa"}
	assert.Contains(t, wantArray, got)

	gotLen := len(builtAllOne.stringMap)
	wantLen := 2
	assert.Equal(t, wantLen, gotLen)

	_, ok := builtAllOne.stringMap["aa"]

	assert.True(t, ok)

	_, ok = builtAllOne.stringMap["Will"]

	assert.True(t, ok)

}
func TestDec_RemovingOneStringWithFrequenceThreeWhenThereAreOtherStringsWithFrequenceOne(t *testing.T) {
	builtAllOne := Constructor()
	builtAllOne.Inc("aa")
	builtAllOne.Inc("aa")
	builtAllOne.Inc("aa")
	builtAllOne.Inc("Will")

	got := builtAllOne.GetMaxKey()
	want := "aa"
	assert.Equal(t, want, got)

	builtAllOne.Dec("aa")
	got = builtAllOne.GetMaxKey()
	want = "aa"
	assert.Equal(t, want, got)

	aaNode := builtAllOne.stringMap["aa"]
	gotFreq := aaNode.frequency
	wantFreq := 2
	assert.Equal(t, wantFreq, gotFreq)

	gotLen := len(builtAllOne.stringMap)
	wantLen := 2
	assert.Equal(t, wantLen, gotLen)

	_, ok := builtAllOne.stringMap["aa"]

	assert.True(t, ok)

	_, ok = builtAllOne.stringMap["Will"]

	assert.True(t, ok)

}

func TestDec_CreatingANewNodeWhenThereIsAGapOf2BetweenItems(t *testing.T) {
	builtAllOne := Constructor()
	builtAllOne.Inc("hello")
	builtAllOne.Inc("goodbye")
	builtAllOne.Inc("farewell")
	builtAllOne.Inc("farewell")
	builtAllOne.Inc("farewell")
	builtAllOne.Inc("hello")
	builtAllOne.Inc("hello")

	maxx1 := builtAllOne.GetMaxKey()
	wantArray := []string{"hello", "farewell"}
	assert.Contains(t, wantArray, maxx1)

	minn1 := builtAllOne.GetMinKey()
	wantArray = []string{"goodbye"}
	assert.Contains(t, wantArray, minn1)

	builtAllOne.Dec("hello")
	maxx1 = builtAllOne.GetMaxKey()
	wantArray = []string{"farewell"}
	assert.Contains(t, wantArray, maxx1)

	minn1 = builtAllOne.GetMinKey()
	wantArray = []string{"goodbye"}
	assert.Contains(t, wantArray, minn1)

}
func TestXxx(t *testing.T) {
	builtAllOne := Constructor()
	builtAllOne.Inc("hello")
	builtAllOne.Inc("goodbye")
	builtAllOne.Inc("hello")
	builtAllOne.Inc("hello")

	maxx1 := builtAllOne.GetMaxKey()
	assert.Equal(t, "hello", maxx1)

	builtAllOne.Inc("leet")
	builtAllOne.Inc("code")
	builtAllOne.Inc("leet")

	builtAllOne.Dec("hello")
	builtAllOne.Inc("leet")
	builtAllOne.Inc("leet")
	builtAllOne.Inc("code")
	builtAllOne.Inc("code")

	maxx2 := builtAllOne.GetMaxKey()
	assert.Equal(t, "leet", maxx2)

}

func TestXxx2(t *testing.T) {
	builtAllOne := Constructor()

	builtAllOne.Inc("a")
	builtAllOne.Inc("b")
	builtAllOne.Inc("c")
	builtAllOne.Inc("d")
	builtAllOne.Inc("e")
	builtAllOne.Inc("f")
	builtAllOne.Inc("g")
	builtAllOne.Inc("h")
	builtAllOne.Inc("i")
	builtAllOne.Inc("j")
	builtAllOne.Inc("k")
	builtAllOne.Inc("l")

	builtAllOne.Dec("a")
	builtAllOne.Dec("b")
	builtAllOne.Dec("c")
	builtAllOne.Dec("d")
	builtAllOne.Dec("e")
	builtAllOne.Dec("f")

	builtAllOne.Inc("g")
	builtAllOne.Inc("h")
	builtAllOne.Inc("i")
	builtAllOne.Inc("j")

	maxx1 := builtAllOne.GetMaxKey()
	wantArray := []string{"g", "h", "i", "j"}
	assert.Contains(t, wantArray, maxx1)

	minn1 := builtAllOne.GetMinKey()
	wantArray = []string{"k", "l"}
	assert.Contains(t, wantArray, minn1)

	builtAllOne.Inc("k")
	builtAllOne.Inc("l")
	maxx1 = builtAllOne.GetMaxKey()
	wantArray = []string{"g", "h", "i", "j", "k", "l"}
	assert.Contains(t, wantArray, maxx1)

	minn1 = builtAllOne.GetMinKey()
	wantArray = []string{"g", "h", "i", "j", "k", "l"}
	assert.Contains(t, wantArray, minn1)

	builtAllOne.Inc("a")
	builtAllOne.Dec("j")
	maxx1 = builtAllOne.GetMaxKey()
	wantArray = []string{"g", "h", "i", "k", "l"}
	assert.Contains(t, wantArray, maxx1)

	minn1 = builtAllOne.GetMinKey()
	wantArray = []string{"a", "j"}
	assert.Contains(t, wantArray, minn1)

}
