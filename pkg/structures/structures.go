package structures

type Params struct {
	Greens      []string
	Yellows     []string
	Greys       []string
	Emoji       []string
	PlayedWords []string
	Turns       int
}

func New() Params {
	return Params{
		Greens:  []string{"_", "_", "_", "_", "_"},
		Yellows: []string{"_", "_", "_", "_", "_"},
		Emoji:   []string{"X", "X", "X", "X", "X"},
	}
}
