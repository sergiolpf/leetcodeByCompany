package mergeintervals

import "sort"

/*
Complexidade onde N eh o numero de intevalos
Tempo: O(N LogN)
Espaco: O(N)
*/
func Merge(intervals [][]int) [][]int {

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	results := [][]int{}
	results = append(results, []int{intervals[0][0], intervals[0][1]})
	indResults := 0

	for i := 1; i < len(intervals); i++ {
		currentStart := intervals[i][0]
		currentEnd := intervals[i][1]

		if currentStart <= results[indResults][1] {
			results[indResults][1] = max(results[indResults][1], currentEnd)
		} else {
			results = append(results, []int{currentStart, currentEnd})
			indResults++
		}

	}

	return results

}
