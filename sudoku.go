package sudoku

type Cell struct {
	value    uint8
	rejected [9]bool
}

type Game struct {
	cells [9 * 9]Cell

	strategies []Strategy
	in         chan *AssertionMessage
}

type AssertionMessage struct {
	positive bool
	index    int // [0, 9*9)
	value    int // [1, 9]
}

func (g *Game) initializeStrategies(initial string) {
	for i, c := range initial {
		value := int(c) - 48
		if value >= 1 && value <= 9 {
			g.in <- &AssertionMessage{true, i, value}
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
		in: make(chan *AssertionMessage, 9*9),
	}

	game.initializeStrategies(initial)

	for m := range game.in {
		if m.positive {

		} else {

		}
	}
}
