package designsnakegame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirections_MovingWithoutFoodSameBodySize(t *testing.T) {
	snakeGame := Constructor(3, 2, [][]int{{-1, -1}})
	snakeGame.Move(Right)
	body := snakeGame.body
	bodyLength := snakeGame.headIndex - snakeGame.tailIndex

	bodyWanted := [][]int{{0, 0}, {0, 1}}
	bodyLengthWanted := 0
	assert.Equal(t, bodyWanted, body, "Corpo moveu com sucesso para a direita")
	assert.Equal(t, bodyLengthWanted, bodyLength, "Corpo permanece com o tamanho desejado")
}
func TestDirections_MovingWithFoodIncreasingBodySize(t *testing.T) {
	snakeGame := Constructor(3, 2, [][]int{{0, 1}})
	snakeGame.Move(Right)
	body := snakeGame.body
	bodyLength := len(snakeGame.body)

	bodyWanted := [][]int{{0, 0}, {0, 1}}
	bodyLengthWanted := 2
	assert.Equal(t, bodyWanted, body, "Corpo moveu com sucesso para a direita e se expandiu por 1")
	assert.Equal(t, bodyLengthWanted, bodyLength, "Corpo cresceu de tamanho por 1.")
}

func TestMove_MovingOutOfBoundsReturningNegativeOne(t *testing.T) {
	snakeGame := Constructor(3, 2, [][]int{{-1, -1}})
	assert.Equal(t, -1, snakeGame.Move(Left))

	snakeGame = Constructor(3, 2, [][]int{{-1, -1}})
	assert.Equal(t, -1, snakeGame.Move(Up))

	snakeGame = Constructor(2, 2, [][]int{{-1, -1}})
	snakeGame.Move(Right)
	snakeGame.Move(Right)
	assert.Equal(t, -1, snakeGame.Move(Right))

	snakeGame = Constructor(2, 2, [][]int{{-1, -1}})
	snakeGame.Move(Down)
	snakeGame.Move(Down)
	assert.Equal(t, -1, snakeGame.Move(Down))
}

func TestDirections_BitingBodyReturningNegativeOne(t *testing.T) {
	snakeGame := Constructor(3, 3, [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 1}})
	snakeGame.Move(Right)
	snakeGame.Move(Right)
	snakeGame.Move(Down)
	snakeGame.Move(Left)
	got := snakeGame.Move(Up)
	want := -1
	assert.Equal(t, want, got, "A cabeca foi para onde o corpo estava, logo, mordeu o proprio corpo e acabou o jogo")
}

func TestDirections_MovingWithoutFoodToTheTailAndNotColiding(t *testing.T) {
	snakeGame := Constructor(3, 3, [][]int{{2, 0}, {0, 0}, {0, 2}, {2, 2}})
	snakeGame.Move(Down)
	snakeGame.Move(Down)
	snakeGame.Move(Right)
	snakeGame.Move(Up)
	snakeGame.Move(Up)
	snakeGame.Move(Left)
	snakeGame.Move(Down)
	snakeGame.Move(Right)
	snakeGame.Move(Right)
	snakeGame.Move(Up)
	snakeGame.Move(Left)
	got := snakeGame.Move(Down)
	want := 3
	assert.Equal(t, want, got, "A cabeca foi para onde o rabo estava, e o rabo deveria mover antes da cabeca colidir")
}

func TestConstructor_InitializesSnakeGameCorrectly(t *testing.T) {
	width := 3
	height := 2
	food := [][]int{{1, 0}, {0, 1}}
	snakeGame := Constructor(width, height, food)

	//verifica se as propriedades foram inicializadas corretamente
	assert.Equal(t, width, snakeGame.width)
	assert.Equal(t, height, snakeGame.height)
	assert.Equal(t, food, snakeGame.foodPositions)
	assert.Equal(t, 0, snakeGame.currentFoodInd)
	assert.Equal(t, 0, snakeGame.headIndex)
	assert.Equal(t, 0, snakeGame.tailIndex)
	assert.Equal(t, [][]int{{0, 0}}, snakeGame.body)
	assert.Equal(t, 0, snakeGame.points)
}

func TestMove_EatingFoodIncreasesPointsAndGrowsSnake(t *testing.T) {
	snakeGame := Constructor(3, 2, [][]int{{1, 0}})

	// Move a cobra para a comida
	points := snakeGame.Move(Down)

	assert.Equal(t, 1, points)              // Verifica se os pontos aumentaram
	assert.Equal(t, 2, len(snakeGame.body)) // Verifica se a cobra cresceu
}

func TestMove_MultipleFoodItemsAreEatenCorrectly(t *testing.T) {
	snakeGame := Constructor(3, 2, [][]int{{1, 0}, {0, 1}, {2, 1}})

	// Come a primeira comida
	snakeGame.Move(Down)
	assert.Equal(t, 1, snakeGame.points)

	// Come a segunda comida
	snakeGame.Move(Right)
	snakeGame.Move(Up)
	assert.Equal(t, 2, snakeGame.points)

	// Come a terceira comida
	snakeGame.Move(Left)
	snakeGame.Move(Down)
	snakeGame.Move(Down)
	snakeGame.Move(Right)
	assert.Equal(t, 3, snakeGame.points)
}
