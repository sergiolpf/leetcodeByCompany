package rankteamsbyvotes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankTeams_EmptyVotes_pass(t *testing.T) {
	got := RankTeams([]string{})

	assert.Equal(t, got, "", "they should be equal")
}
func TestRankTeams_pass(t *testing.T) {
	got := RankTeams([]string{"ABC", "ACB", "ABC", "ACB", "ACB"})

	assert.Equal(t, got, "ACB", "they should be equal")
}

func TestRankTeams_drawByFirstPositions_pass(t *testing.T) {
	got := RankTeams([]string{"ABC", "ABC", "BAC", "BCA"})

	assert.Equal(t, got, "BAC", "they should be equal")
}
