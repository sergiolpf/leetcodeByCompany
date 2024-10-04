package maximumsquareareabyremovingfencesfromafield

import (
	"math"
	"sort"
)

const MOD = 1e9 + 7

func MaximizeSquareArea(m, n int, hFences, vFences []int) int {
	// Adicionando as cercas fixas ao redor do campo
	hFences = append([]int{1}, append(hFences, m)...)
	vFences = append([]int{1}, append(vFences, n)...)

	// Ordenando as cercas
	sort.Ints(hFences)
	sort.Ints(vFences)

	gap := make(map[int]bool)

	// Encontrando o maior espaço entre cercas horizontais
	for i := 0; i < len(hFences); i++ {
		for j := i + 1; j < len(hFences); j++ {
			gap[hFences[j]-hFences[i]] = true
		}
	}

	maxArea := math.MinInt64
	for i := 0; i < len(vFences); i++ {
		for j := i + 1; j < len(vFences); j++ {
			if gap[vFences[j]-vFences[i]] {
				maxArea = max(maxArea, (vFences[j]-vFences[i])*(vFences[j]-vFences[i]))
			}
		}
	}

	if maxArea < 0 {
		return -1
	}
	// Calcular a área máxima do quadrado, considerando o módulo
	return int(maxArea % MOD)
}
