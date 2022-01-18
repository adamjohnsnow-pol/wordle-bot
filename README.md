# WordleBot
See whether a simple looping algorithm would guess the answer before you did

### To play
```
go mod tidy
go run .
```
WordleBot can start with a guess of your choosing, it will ask for that first if you want to use one, or hit enter for it to start with a random guess.
Then WordleBot will ask for the word it is trying to find, type it in (don't worry, it won't cheat)
It will then show you it's results - the grid, turns taken (if <=6) and the words used.
```
WordleBot: Want to give me my first guess? (Hit enter if not)
-> guess
WordleBot: Tell me the answer first
-> rebus
WordleBot got:
⬜🟨🟨⬜🟩
⬜⬜🟨🟨🟩
🟨🟨🟨⬜🟩
🟩🟩🟨🟩🟩
🟩🟩🟩🟩🟩
5 / 6
[GUESS CLUES UREAS RESUS REBUS]
```
#### Todo
-Not sure if the yellow/grey logic is correct for duplicated letters
