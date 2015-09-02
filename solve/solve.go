package main

import "github.com/jneu/sudoku"

var games = [...]string{
	"5..4.6..89.8.325.41.........4.57.3.1....2....8.3.41.2.........23.921.7.66..7.3..9",
	".1...54765.43.7...72...4..58.7.3...9.3.....4.2...4.8.73..4...91...1.37.81657...4.",
}

func main() {
	for _, initial := range games {
		sudoku.Solve(initial)
	}
}
