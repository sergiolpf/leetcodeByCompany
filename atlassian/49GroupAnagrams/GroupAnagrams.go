package groupanagrams

import "sort"

/*
Complexidade
Tempo: O(N * MLogM) - onde N eh o tamanho do input e M eh a media do tamanho da string
Espaco: O (M * N)
*/
func groupAnagrams(strs []string) [][]string {

	if len(strs) == 1 {
		return [][]string{strs}
	}

	results := [][]string{}

	orderedMap := map[string][]string{}

	for _, str := range strs {
		s_anagram := []rune(str)
		sort.Slice(s_anagram, func(i, j int) bool {
			return s_anagram[i] < s_anagram[j]
		})

		orderedMap[string(s_anagram)] = append(orderedMap[string(s_anagram)], str)
	}

	for _, elements := range orderedMap {
		results = append(results, elements)
	}

	return results
}
