package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"wordle-bot/pkg/game"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("WordleBot: Want to give me my first guess? (Hit enter if not)")
	fmt.Print("-> ")
	initial, _ := reader.ReadString('\n')
	initial = strings.Replace(initial, "\n", "", -1)
	initial = strings.ToUpper(initial)

	fmt.Println("WordleBot: Tell me the answer first")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		fmt.Println("WordleBot got:")
		res := game.PlayGame(strings.ToUpper(text), initial)
		if strings.Join(res.Emoji, "") != "GGGGG" {
			fmt.Println("X / 6")
		} else {
			fmt.Println(len(res.PlayedWords), "/ 6")
		}
		fmt.Println(res.PlayedWords)
	}
}
