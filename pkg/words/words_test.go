package words

import (
	"testing"
	"wordle-bot/pkg/structures"

	"github.com/stretchr/testify/assert"
)

func TestGreens(t *testing.T) {
	p := structures.Params{
		Greens: []string{"A", "_", "_", "_", "E"},
	}
	assert.True(t, hasAllGreens("ALIVE", p))
	assert.False(t, hasAllGreens("BLIVE", p))
	assert.False(t, hasAllGreens("ALIVB", p))
}

func TestInclusion(t *testing.T) {
	p := structures.Params{
		Yellows: []string{"A", "B", "C", "_", "_"},
	}

	assert.True(t, hasYellows("XABCX", p))
	assert.False(t, hasYellows("ABCXX", p))
	assert.False(t, hasYellows("AXXXX", p))
	assert.True(t, hasYellows("XXABC", p))
	assert.True(t, hasYellows("XCXBA", p))
}

func TestExclusion(t *testing.T) {
	p := structures.Params{
		Greys: []string{"A", "B", "C"},
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
		Greens: []string{"A", "_", "_", "_", "E"},
	}
	candidates := findAllCandidates(p)

	assert.True(t, Contains("ABELE", candidates))
	assert.False(t, Contains("SLIVE", candidates))
	assert.True(t, Contains("AGLEE", candidates))
	assert.True(t, Contains("AGAVE", candidates))
	assert.False(t, Contains("AGAIN", candidates))
}

func TestPickRandomWord(t *testing.T) {
	words := []string{"THIS", "THAT", "OTHER"}
	randomChoice := pickRandomWord(words)
	assert.True(t, Contains(randomChoice, words))
}
