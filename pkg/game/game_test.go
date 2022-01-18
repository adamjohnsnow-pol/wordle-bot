package game

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	fmt.Println("WordleBot got:")
	res := PlayGame("CRANK")
	if strings.Join(res.Emoji, "") != "GGGGG" {
		fmt.Println("X / 6")
	} else {
		fmt.Println(len(res.PlayedWords), "/ 6")
	}

	fmt.Println(res.PlayedWords)

}

func TestEmoji(t *testing.T) {
	res := printEmoji([]string{"G", "Y", "X"})
	assert.Equal(t, "🟩🟨⬜", res)
}
