package main

type Token int

const (
	GAME_STARTED = iota
	GAME_SHUTDOWN
	KILL
	CAUSE_BY
)

var tokens = []string{
	GAME_STARTED:  "InitGame:",
	GAME_SHUTDOWN: "ShutdownGame:",
	KILL:          "killed",
	CAUSE_BY:      "by",
}

func (t Token) String() string {
	return tokens[t]
}
