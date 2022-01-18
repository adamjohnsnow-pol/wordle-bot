package game

import (
	"fmt"
	"strings"
	"wordle-bot/pkg/logic"
	"wordle-bot/pkg/structures"
	"wordle-bot/pkg/words"
)

func PlayGame(target string) structures.Params {
	result := structures.New()
	for {
		result.Turns++
		guess := words.GetAWord(&result)
		if guess == "" {
			break
		}
		logic.CompareLetters(guess, target, &result)
		fmt.Println(printEmoji(result.Emoji))
		if !words.Contains("_", result.Greens) || result.Turns == 6 {
			break
		}
	}
	return result
}

func printEmoji(emoji []string) string {
	withGreen := strings.ReplaceAll(strings.Join(emoji, ""), "G", "🟩")
	withYellow := strings.ReplaceAll(withGreen, "Y", "🟨")
	withGrey := strings.ReplaceAll(withYellow, "X", "⬜")
	return withGrey
}
