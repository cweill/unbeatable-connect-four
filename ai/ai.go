package ai

import (
	"time"
	"c4/game"
)

type AI struct {
	Player game.Player
}

func (a *AI) ChooseMove(s *game.State) game.Column {
	time.Sleep(500 * time.Millisecond)
	return a.bestMove(s, 5)
}

type value int

func (a *AI) bestMove(s *game.State, depth int) game.Column {
	bestCol := game.Column(-1)
	bestVal := value(-1)
	for i := game.Column(0); i < game.Column(len(s.Grid)); i++ {
		ss, _ := s.Move(i)
		if a.minmax(ss, depth, a.Player) > bestVal {
			bestCol = i
		}
	}
	return bestCol
}

func (a *AI) minmax(s *game.State, depth int, p game.Player) value {
	if depth == 0 || s.IsGameOver() {
		return a.stateValue(s)
	}
	bestVal := value(1)
	if p == a.Player {
		bestVal = value(-1)
	}
	for i := game.Column(0); i < game.Column(len(s.Grid)); i++ {
		ss, _ := s.Move(i)
		val := a.minmax(ss, depth - 1, p.Opponent())
		if p == a.Player {
			if val >= bestVal {
				bestVal = val
			}
		} else {
			if val <= bestVal {
				bestVal = val
			}
		}
	}
	return bestVal
}

func (a *AI) stateValue(s *game.State) value {
	if !s.IsGameOver() {
		return 0
	}
	if s.Turn == a.Player {
		return 1
	}
	return -1
}