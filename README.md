# WordleBot
See whether a simple looping algorithm would guess the answer before you did

### To play
```
go mod tidy
go run .
```
WordleBot will ask for the word it is trying to find, type it in (don't worry, it won't cheat)
It will then show you it's results - the grid, turns taken (if <=6) and the words used.
```
WordleBot: Tell me the answer first
-> REBUS
⬜⬜🟨⬜⬜
⬜🟩⬜⬜🟨
⬜🟩⬜🟨🟨
🟩🟩⬜🟨⬜
🟩🟩🟩🟨⬜
🟩🟩🟩🟨⬜
WordleBot got:
X / 6
```
#### Todo
-Not sure if the yellow/grey logic is correct for duplicated letters
