package words

import (
	"testing"
	"wordle-bot/pkg/structures"

	"github.com/stretchr/testify/assert"
)

func contains(target string, slice []string) bool {
	for _, word := range slice {
		if word == target {
			return true
		}
	}
	return false
}

func TestGreens(t *testing.T) {
	p := structures.Params{
		LetterOne:  "A",
		LetterTwo:  "",
		LetterFive: "E",
	}
	assert.True(t, hasAllGreens("ALIVE", p))
	assert.False(t, hasAllGreens("BLIVE", p))
	assert.False(t, hasAllGreens("ALIVB", p))
}

func TestInclusion(t *testing.T) {
	p := structures.Params{
		Includes: []string{"A", "B", "C"},
	}

	assert.True(t, hasYellows("ABCXX", p))
	assert.False(t, hasYellows("ABXXX", p))
	assert.False(t, hasYellows("AXXXX", p))
	assert.True(t, hasYellows("XXABC", p))
	assert.True(t, hasYellows("XCXBA", p))
}

func TestExclusion(t *testing.T) {
	p := structures.Params{
		Excludes: []string{"A", "B", "C"},
	}

	assert.False(t, hasNoGreys("ABCXX", p))
	assert.False(t, hasNoGreys("ABXXX", p))
	assert.False(t, hasNoGreys("AXXXX", p))
	assert.False(t, hasNoGreys("XXABC", p))
	assert.False(t, hasNoGreys("XCXBA", p))
	assert.True(t, hasNoGreys("XXXXX", p))
}

func TestFindCandidates(t *testing.T) {
	p := structures.Params{
		LetterOne:  "A",
		LetterTwo:  "",
		LetterFive: "E",
	}
	candidates := findAllCandidates(p)

	assert.True(t, contains("ABELE", candidates))
	assert.False(t, contains("SLIVE", candidates))
	assert.True(t, contains("AGLEE", candidates))
	assert.True(t, contains("AGAVE", candidates))
	assert.False(t, contains("AGAIN", candidates))
}

func TestPickRandomWord(t *testing.T) {
	words := []string{"THIS", "THAT", "OTHER"}
	randomChoice := pickRandomWord(words)
	assert.True(t, contains(randomChoice, words))
}
