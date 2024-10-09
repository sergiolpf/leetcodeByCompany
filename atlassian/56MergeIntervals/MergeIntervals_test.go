package mergeintervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge_EmtpyInput(t *testing.T) {
	got := Merge([][]int{})
	want := [][]int{}

	assert.Equal(t, want, got, "Deveria retornar um slice vazio para input vazio")
}
func TestMerge_SingleInterval(t *testing.T) {
	got := Merge([][]int{
		{1, 3},
	})

	want := [][]int{
		{1, 3},
	}

	assert.Equal(t, want, got, "Deveria retornar o mesmo intervalo quando ha somente 1 intervalo como entrada")
}
func TestMerge_OverlappingIntervals(t *testing.T) {
	got := Merge([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	})

	want := [][]int{
		{1, 6},
		{8, 10},
		{15, 18},
	}

	assert.Equal(t, want, got, "Deveria mesclar intervalos sobrepostos corretamente")
}

func TestMerge_NonOverlappingIntervals(t *testing.T) {
	got := Merge([][]int{
		{1, 3},
		{8, 10},
		{15, 18},
	})

	want := [][]int{
		{1, 3},
		{8, 10},
		{15, 18},
	}

	assert.Equal(t, want, got, "Deveria retornar os mesmos intervalos ja que nao ha sobreposicao")
}

func TestMerge_IntervalTouchingEdges(t *testing.T) {
	got := Merge([][]int{
		{1, 4},
		{4, 7},
		{8, 10},
	})

	want := [][]int{
		{1, 7},
		{8, 10},
	}

	assert.Equal(t, want, got, "Deveria mesclar intervalos que se tocam nas bordas")
}

func TestMerge_UnsortedIntervalszx(t *testing.T) {
	got := Merge([][]int{
		{4, 7},
		{8, 10},
		{1, 4},
	})

	want := [][]int{
		{1, 7},
		{8, 10},
	}

	assert.Equal(t, want, got, "Deveria mesclar corretamente mesmo ")
}
