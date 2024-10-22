package longeststringchain

import "sort"

/*
Complexidade:
Tempo:  O(N log N + N*M + L) - onde N eh o numero de palavras noa rray
M eh o tamanho medio das palavras no array words
L o tamanho de cada cadeia encontrada.const

Espaco: O(N*M)
*/
func LongestStrChain(words []string) int {

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	mapOfWords := make(map[string]int)
	maxChain := 0

	for _, word := range words {
		mapOfWords[word] = 1
		for i := 0; i < len(word); i++ {
			prevWord := word[:i] + word[i+1:]
			if val, ok := mapOfWords[prevWord]; ok {
				mapOfWords[word] = max(mapOfWords[word], val+1)

			}
		}
		maxChain = max(maxChain, mapOfWords[word])
	}

	return maxChain
}
