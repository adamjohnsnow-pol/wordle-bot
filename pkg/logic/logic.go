package logic

import (
	"strings"
	"wordle-bot/pkg/structures"
	"wordle-bot/pkg/words"
)

func CompareLetters(guess, target string, result *structures.Params) {
	compareGreens(guess, target, result)
	compareYellows(guess, target, result)
}
func compareGreens(guess, target string, result *structures.Params) {
	guessLetters := strings.Split(guess, "")
	targetLetters := strings.Split(target, "")
	for i, letter := range guessLetters {
		if letter == targetLetters[i] {
			result.Greens[i] = letter
			result.Emoji[i] = "G"
		}
	}
}

func compareYellows(guess, target string, result *structures.Params) {
	guessLetters := strings.Split(guess, "")
	targetLetters := strings.Split(target, "")
	for i, letter := range guessLetters {
		if words.Contains(letter, targetLetters) && !words.Contains(letter, result.Greens) {
			result.Yellows[i] = letter
			result.Emoji[i] = "Y"
		}
		if !words.Contains(letter, targetLetters) && !words.Contains(letter, result.Greens) && !words.Contains(letter, result.Greys) {
			result.Greys = append(result.Greys, letter)
			result.Emoji[i] = "X"
		}
	}
}
