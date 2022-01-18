package logic

import (
	"testing"
	"wordle-bot/pkg/structures"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	guess  string
	target string
	result []string
}

func TestCompareGreens(t *testing.T) {
	testCases := []testCase{
		{"AXXXX", "AZZZZ", []string{"A", "_", "_", "_", "_"}},
		{"XAXXX", "ZAZZZ", []string{"_", "A", "_", "_", "_"}},
		{"XXAXX", "ZZAZZ", []string{"_", "_", "A", "_", "_"}},
		{"XXXAX", "ZZZAZ", []string{"_", "_", "_", "A", "_"}},
		{"XXXXA", "ZZZZA", []string{"_", "_", "_", "_", "A"}},
		{"AXXXA", "AZZZA", []string{"A", "_", "_", "_", "A"}},
	}
	for _, tCase := range testCases {
		result := structures.New()
		compareGreens(tCase.guess, tCase.target, &result)
		assert.Equal(t, tCase.result, result.Greens)
	}
}
func TestCompareEmoji(t *testing.T) {
	testCases := []testCase{
		{"AXXXX", "AZZZZ", []string{"G", "X", "X", "X", "X"}},
		{"XAXXX", "ZAZZZ", []string{"X", "G", "X", "X", "X"}},
		{"XXAXX", "ZZAZZ", []string{"X", "X", "G", "X", "X"}},
		{"XXXAX", "ZZZAZ", []string{"X", "X", "X", "G", "X"}},
		{"XXXXA", "ZZZZA", []string{"X", "X", "X", "X", "G"}},
		{"AXXXA", "AZZZA", []string{"G", "X", "X", "X", "G"}},
	}
	for _, tCase := range testCases {
		result := structures.New()
		compareGreens(tCase.guess, tCase.target, &result)
		assert.Equal(t, tCase.result, result.Emoji)
	}
}
func TestCompareYellows(t *testing.T) {
	testCases := []testCase{
		{"AXXXX", "ZAZZZ", []string{"A", "_", "_", "_", "_"}},
		{"ABXXX", "ZAZBZ", []string{"A", "B", "_", "_", "_"}},
	}
	for _, tCase := range testCases {
		result := structures.New()
		compareYellows(tCase.guess, tCase.target, &result)
		assert.Equal(t, tCase.result, result.Yellows)
	}
}
func TestCompareYellowEmoji(t *testing.T) {
	testCases := []testCase{
		{"AXXXX", "ZAZZZ", []string{"Y", "X", "X", "X", "X"}},
		{"ABXXX", "ZAZBZ", []string{"Y", "Y", "X", "X", "X"}},
	}
	for _, tCase := range testCases {
		result := structures.New()
		compareYellows(tCase.guess, tCase.target, &result)
		assert.Equal(t, tCase.result, result.Emoji)
	}
}
func TestGreenAndYellow(t *testing.T) {
	result := structures.New()
	compareGreens("ABACR", "AXFZC", &result)
	compareYellows("ABACR", "AXFZC", &result)
	expect := structures.Params{
		Greens:  []string{"A", "_", "_", "_", "_"},
		Yellows: []string{"_", "_", "_", "C", "_"},
		Greys:   []string{"B", "R"},
		Emoji:   []string{"G", "X", "X", "Y", "X"},
	}
	assert.Equal(t, expect, result)
	compareYellows("ACALR", "AXFZC", &result)
	expect.Yellows = []string{"_", "C", "_", "C", "_"}
	assert.Equal(t, expect.Yellows, result.Yellows)
}
