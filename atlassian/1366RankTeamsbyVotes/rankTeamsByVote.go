package rankteamsbyvotes

import "sort"

type TeamVotes struct {
	team  byte
	votes []int
}

/*
Complexidade:
Tempo: O(M * N * LogN) - N - Numero de times e M - numero de votos
Espaco: O (M * N)
*/
func RankTeams(votes []string) string {

	results := []TeamVotes{}

	voteResultsMap := make(map[byte][]int, len(votes))
	for _, vote := range votes {
		for i := 0; i < len(vote); i++ {
			if _, ok := voteResultsMap[vote[i]]; !ok {
				voteResultsMap[vote[i]] = make([]int, len(vote))
			}
			voteResultsMap[vote[i]][i]++
		}
	}

	for team, teamResults := range voteResultsMap {
		results = append(results, TeamVotes{
			team:  team,
			votes: teamResults})
	}

	sort.Slice(results, func(i, j int) bool {
		for ind := 0; ind < len(results[i].votes); ind++ {
			if results[i].votes[ind] != results[j].votes[ind] {
				return results[i].votes[ind] > results[j].votes[ind]
			}
		}
		return results[i].team < results[j].team
	})

	var finalResult string
	for _, v := range results {
		finalResult += string(v.team)
	}

	return finalResult
}
