// Package ai contains logic for an artificial intelligence player.
package ai

import (
	"c4/game"
)

const maxDepth = 5

// AI represents an artificial intelligence player.
type AI struct {
	Player game.Player
}

// ChooseMove will return the best possible move given the state of the game.
func (a *AI) ChooseMove(s *game.State) game.Column {
	col, _ := a.minmax(s, maxDepth, true)
	return col
}

// value represents the utility of a move for Min-Max.
type value int

const (
	infinite = value(99999999)
	uncertain = value(0)
	win = value(10)
	loss = value(-10)
)

// stateValue assigns a value to a state. A losing state gets a negative score,
// and a winning state gets a positive score.
func (a *AI) stateValue(s *game.State, depth int) value {
	if !s.IsGameOver() {
		return uncertain
	}
	// We want to penalize losing and reward winning more at a shallower depth.
	multi := value(depth + 1)
	if s.Turn != a.Player {
		// If the game's turn is over and it's not the AI player's turn, it 
		// means the AI was the last to make a move.
		return win * multi
	}
	return loss * multi
}

// minmax implements the Min-Max algorithm and returns the best move and value
// of that move.
func (a *AI) minmax(s *game.State, depth int, maxPlayer bool) (game.Column, value) {
	if depth == 0 || s.IsGameOver() {
		return game.MaxColumn, a.stateValue(s, depth)
	}
	bestVal := infinite
	if maxPlayer {
		bestVal = -infinite
	}
	var bestCol game.Column
	for i := game.Column(0); i <= game.MaxColumn; i++ {
		if !s.IsValidMove(i) {
			continue
		}
		ss, _ := s.Move(i)
		_, val := a.minmax(ss.NextTurn(), depth - 1, !maxPlayer)
		if maxPlayer {
			if val > bestVal {
				bestVal = val
				bestCol = i
			}
		} else {
			if val < bestVal {
				bestVal = val
				bestCol = i
			}
		}
	}
	return bestCol, bestVal
}
