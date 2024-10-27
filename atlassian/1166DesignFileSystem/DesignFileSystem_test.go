package designfilesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	 func TestConstructor_basicConstruction(t *testing.T) {
		fileSystem := Constructor()

		assert.Empty(t, fileSystem.head.children)
		assert.Equal(t, 1, len(fileSystem.validPaths))
		//assert.Empty(t, fileSystem.validPaths)

}
*/
func TestCreatePath_InvalidPaths(t *testing.T) {
	fileSystem := Constructor()

	assert.False(t, fileSystem.CreatePath("", 1))
	assert.False(t, fileSystem.CreatePath("/", 2))
}

func TestCreatePath_CreatingRepeatedPaths(t *testing.T) {
	fileSystem := Constructor()

	fileSystem.CreatePath("/leetcode", 1)
	assert.False(t, fileSystem.CreatePath("/leetcode", 2))

}
func TestCreatePath_CreatingPathsWithNoParents(t *testing.T) {
	fileSystem := Constructor()

	assert.False(t, fileSystem.CreatePath("/leetcode/folder1", 1))

}

func TestCreatePath_ValidSinglePaths(t *testing.T) {
	fileSystem := Constructor()

	assert.True(t, fileSystem.CreatePath("/leetcode", 1))
	assert.Equal(t, 1, fileSystem.Get("/leetcode"))

}

func TestCreatePath_ValidSingleSubPath(t *testing.T) {
	fileSystem := Constructor()

	fileSystem.CreatePath("/leetcode", 1)
	assert.True(t, fileSystem.CreatePath("/leetcode/problem", 2))

	assert.Equal(t, 2, fileSystem.Get("/leetcode/problem"))

}

func TestCreatePath_ValidMultipleSubPath(t *testing.T) {
	fileSystem := Constructor()

	fileSystem.CreatePath("/leetcode", 1)
	assert.True(t, fileSystem.CreatePath("/leetcode/example1", 2))
	assert.True(t, fileSystem.CreatePath("/leetcode/example2", 3))

	assert.Equal(t, 2, fileSystem.Get("/leetcode/example1"))
	assert.Equal(t, 3, fileSystem.Get("/leetcode/example2"))

}

func TestGet_InvalidPaths(t *testing.T) {
	fileSystem := Constructor()

	fileSystem.CreatePath("", 1)
	fileSystem.CreatePath("/", 2)

	assert.Equal(t, -1, fileSystem.Get(""))
	assert.Equal(t, -1, fileSystem.Get("/"))
	assert.Equal(t, -1, fileSystem.Get("/pathNotCreated"))

}
