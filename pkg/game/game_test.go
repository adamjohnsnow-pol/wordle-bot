package game

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	fmt.Println("WordleBot got:")
	res := PlayGame("CRANK", "")
	if strings.Join(res.Emoji, "") != "GGGGG" {
		fmt.Println("X/ 6")
		assert.NotEqual(t, "CRANK", res.PlayedWords[5])
	} else {
		fmt.Println(len(res.PlayedWords), "/ 6")
		assert.Equal(t, "CRANK", res.PlayedWords[len(res.PlayedWords)-1])
	}
	fmt.Println(res.PlayedWords)
	//NO ACTUAL TEST
}
func TestGameWithStartWord(t *testing.T) {
	res := PlayGame("CRANK", "AUDIO")
	assert.Equal(t, "AUDIO", res.PlayedWords[0])
}
func TestEmoji(t *testing.T) {
	res := printEmoji([]string{"G", "Y", "X"})
	assert.Equal(t, "🟩🟨⬜", res)
}
