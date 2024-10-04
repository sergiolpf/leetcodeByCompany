package maximumareaoflongestdiagonalrectangle

import (
	"math"
)

func AreaOfMaxDiagonal(dimensions [][]int) int {

	maxDiagonal := 0.0
	maxArea := 0

	for i := 0; i < len(dimensions); i++ {
		length := float64(dimensions[i][0])
		width := float64(dimensions[i][1])
		diagonal := math.Sqrt(width*width + length*length)
		if diagonal > maxDiagonal {
			maxDiagonal = diagonal
			maxArea = int(width * length)
		} else if diagonal == maxDiagonal {
			maxArea = max(maxArea, int(length*width))
		}
	}

	return maxArea
}
