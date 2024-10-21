package wordsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExist_WordOnTopRightCorner(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	got := Exist(board, "SEE")
	assert.True(t, got)
}

func TestExist_WordOnBottomLeftCorner(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	got := Exist(board, "ADE")
	assert.True(t, got)
}

func TestExist_SingleLetterWord(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	got := Exist(board, "B")
	assert.True(t, got)
}

func TestExist_EmptyBoard(t *testing.T) {
	board := [][]byte{}

	got := Exist(board, "ABCCED")
	assert.False(t, got)
}

func TestExist_EmptyWord(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	got := Exist(board, "")
	assert.True(t, got)
}

func TestExist_OneDimensionBoardWithOneCharacter(t *testing.T) {
	board := [][]byte{{'A'}}
	got := Exist(board, "A")
	assert.True(t, got)
}

func TestExist_OneDimensionBoardWithDifferentCharacter(t *testing.T) {
	board := [][]byte{{'A'}}
	got := Exist(board, "B")
	assert.False(t, got)
}

func TestExist_WordBiggerThanBoard(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	got := Exist(board, "ABCCEDASDFASDFASDFASDFASDF")
	assert.False(t, got)
}

func TestExist_WordAppearsMultipleTimes(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'A'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'A'},
	}

	got := Exist(board, "ASA")
	assert.True(t, got)
}
