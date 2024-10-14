package designsnakegame

import "reflect"

type SnakeGame struct {
	width, height  int
	foodPositions  [][]int
	currentFoodInd int
	headIndex      int
	tailIndex      int
	body           [][]int
	points         int
	bodyMap        map[[2]int]bool
}

const (
	Right string = "R"
	Left  string = "L"
	Up    string = "U"
	Down  string = "D"
)

func Constructor(width int, height int, food [][]int) SnakeGame {
	return SnakeGame{
		width:          width,
		height:         height,
		foodPositions:  food,
		currentFoodInd: 0,
		headIndex:      0,
		tailIndex:      0,
		body:           [][]int{{0, 0}},
		points:         0,
		bodyMap:        make(map[[2]int]bool),
	}

}

func (this *SnakeGame) Move(direction string) int {

	headRow := this.body[this.headIndex][0]
	headCol := this.body[this.headIndex][1]
	switch direction {
	case Right:
		headCol++
	case Left:
		headCol--
	case Up:
		headRow--
	case Down:
		headRow++
	}
	return this.ValidatePosition([]int{headRow, headCol})
}

/*
Complexidade
Tempo: Parece absurdo mas eu acho que o tempo pode ser no maximo O(1)
Espaco: O(N) - onde N eh o tamanho da cobra.
*/
func (this *SnakeGame) ValidatePosition(position []int) int {
	if position[0] < 0 || position[0] >= this.width ||
		position[1] < 0 || position[1] >= this.height {
		return -1
	}

	if this.bodyMap[[2]int(position)] &&
		!reflect.DeepEqual(this.body[this.tailIndex], position) {
		return -1
	}

	if this.currentFoodInd < len(this.foodPositions) &&
		reflect.DeepEqual(position, this.foodPositions[this.currentFoodInd]) {

		this.points++
		if len(this.foodPositions) > 0 {
			this.currentFoodInd++
		}
	} else {
		delete(this.bodyMap, [2]int{this.body[this.tailIndex][0], this.body[this.tailIndex][1]})
		this.tailIndex++
	}

	this.body = append(this.body, []int{position[0], position[1]})
	this.headIndex++
	//this.tail = [2]int{this.body[0][0], this.body[0][1]}
	this.bodyMap[[2]int(position)] = true

	return this.points
}

/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
