package highaccessemployees

import (
	"sort"
	"strconv"
)

const (
	ONE_HOUR    = 100
	WINDOW_SIZE = 2
)

/*
Complexidade
Tempo: O(NlogN)
Espaco: O(n)
*/
func findHighAccessEmployees(access_times [][]string) []string {

	response := []string{}
	mapOfNameAndTimes := map[string][]int{}
	for index := 0; index < len(access_times); index++ {

		timeInString, _ := strconv.Atoi(access_times[index][1])
		mapOfNameAndTimes[access_times[index][0]] =
			append(mapOfNameAndTimes[access_times[index][0]], timeInString)
	}

	for _, accesstimes := range mapOfNameAndTimes {
		sort.Slice(accesstimes, func(i int, j int) bool {
			return accesstimes[i] < accesstimes[j]
		})
	}

	var start, end int
	for name, orderedAccessTimes := range mapOfNameAndTimes {

		start, end = 0, WINDOW_SIZE
		for end < len(orderedAccessTimes) {

			if orderedAccessTimes[end]-orderedAccessTimes[start] < ONE_HOUR {
				response = append(response, name)
				break
			} else {
				start++
				end++
			}
		}

	}
	return response

}
