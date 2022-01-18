package game

import (
	"fmt"
	"strings"
	"wordle-bot/pkg/logic"
	"wordle-bot/pkg/structures"
	"wordle-bot/pkg/words"
)

func PlayGame(target, initial string) structures.Params {
	result := structures.New()
	var guess string
	if initial != "" {
		guess = initial
	} else {
		guess = words.GetAWord(&result)
	}

	for {
		result.Turns++
		if guess == "" {
			break
		}
		logic.CompareLetters(guess, target, &result)
		result.PlayedWords = append(result.PlayedWords, guess)
		fmt.Println(printEmoji(result.Emoji))
		if !words.Contains("_", result.Greens) || result.Turns == 6 {
			break
		}
		guess = words.GetAWord(&result)
	}
	return result
}

func printEmoji(emoji []string) string {
	withGreen := strings.ReplaceAll(strings.Join(emoji, ""), "G", "🟩")
	withYellow := strings.ReplaceAll(withGreen, "Y", "🟨")
	withGrey := strings.ReplaceAll(withYellow, "X", "⬜")
	return withGrey
}
