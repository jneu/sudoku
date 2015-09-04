package sudoku

import "github.com/jneu/sudoku/common"
import "github.com/jneu/sudoku/strategy"

type Cell struct {
	value    uint8
	rejected [9]bool
}

type Game struct {
	cells [9 * 9]Cell

	strategies []strategy.Strategy
	in         chan *common.Assertion
}

func (g *Game) initializeStrategies(initial string) {
	for i, c := range initial {
		value := int(c) - 48
		if value >= 1 && value <= 9 {
			g.in <- &common.Assertion{value, i}
		} else {
			g.strategies = append(g.strategies, NewCellStrategy(i, g.in))
		}
	}
}

func Solve(initial string) {
	game := &Game{
		// This buffer should be large enough to handle at least one positive
		// assertion for each cell. These are generated when we initialize our
		// strategies and we don't want to block when we add them.
		in: make(chan *common.Assertion, 9*9),
	}

	game.initializeStrategies(initial)

	for m := range game.in {
		if m.positive {

		} else {

		}
	}
}
