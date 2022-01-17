package logic

import (
	"testing"
	"wordle-bot/pkg/structures"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	result := compareWords("AXXXX", "XAXXX")
	expect := structures.Params{Includes: []string{"A"}, Excludes: []string{"X"}}

	assert.Equal(t, expect, result)
}
