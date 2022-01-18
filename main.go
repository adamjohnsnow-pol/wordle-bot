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
	fmt.Println("WordleBot: Tell me the answer first")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		res := game.PlayGame(strings.ToUpper(text))
		fmt.Println("WordleBot got:")
		if strings.Join(res.Emoji, "") != "GGGGG" {
			fmt.Println("X / 6")
		} else {
			fmt.Println(len(res.PlayedWords), "/ 6")
		}
		fmt.Println(res.PlayedWords)
	}
}
