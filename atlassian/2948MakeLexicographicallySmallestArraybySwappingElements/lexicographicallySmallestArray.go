package lexicographicallySmallestArray

import "sort"

func lexicographicallySmallestArray(nums []int, limit int) []int {
	size := len(nums)
	result := make([]int, size)
	sortedNums := make([]int, size)

	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	//esse mapa, vai dizer em que grupo cada elemento do sortedNums array estara
	groups := make(map[int]int)
	//O mapa groupstarter vai ter os mesmos grupos, so que dessa vez teremos
	//o indice dentro do array ordenado onde o grupo comecara.
	groupStarter := make(map[int]int)
	groupsCount := 0
	for i := 0; i < size; i++ {
		if i > 0 && sortedNums[i]-sortedNums[i-1] > limit {
			groupsCount++
		}
		groups[sortedNums[i]] = groupsCount
		//isso vai dizer qual o indice dentro do array ordenado que comeca
		// o proximo grupo.
		if _, ok := groupStarter[groupsCount]; !ok {
			groupStarter[groupsCount] = i
		}
	}

	for index := 0; index < size; index++ {
		groupOfNumber := groups[nums[index]]
		sortedNumsIndex := groupStarter[groupOfNumber]
		result[index] = sortedNums[sortedNumsIndex]
		groupStarter[groupOfNumber]++

	}

	return result
}
