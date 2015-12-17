// Package ai contains logic for an artificial intelligence player.
package ai

import (
	"../game"
	"math/rand"
)

// Difficulty represents how challenging the AI will be.
type Difficulty int

const (
	// Easy is beatable.
	Easy Difficulty = iota
	// Medium is doable.
	Medium
	// Hard is a challenge.
	Hard
	// Impossible.
	Impossible
)

// How many moves the AI should look ahead.
const (
	easyDepth = 1
	mediumDepth = 3
	hardDepth = 5
	impossibleDepth = 10
)

// AI represents an artificial intelligence player.
type AI struct {
	Player game.Player
	Difficulty Difficulty
}

// ChooseMove will return the best possible move given the state of the game, 
// and the difficulty of the AI.
func (a *AI) ChooseMove(s *game.State) game.Column {
	var col game.Column
	switch a.Difficulty {
	case Easy:
		col, _ = a.minmax(s, easyDepth, true)
	case Medium:
		col, _ = a.minmax(s, mediumDepth, true)
	case Hard:
		col, _ = a.minmax(s, hardDepth, true)
	case Impossible:
		col, _ = a.alphabeta(s, impossibleDepth, -infinite, infinite, true)
	}
	return col
}

// value represents the utility of a move for Min-Max.
type value int

const (
	infinite  = value(99999999)
	uncertain = value(0)
	win       = value(10)
	loss      = value(-10)
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
	var bestCol game.Column
	if maxPlayer {
		v := -infinite
		for _, i := range shuffledColumns() {
			if !s.IsValidMove(i) {
				continue
			}
			ss, _ := s.Move(i)
			if _, val := a.minmax(ss.NextTurn(), depth-1, !maxPlayer); val > v {
				v = val
				bestCol = i
			}
		}
		return bestCol, v
	}
	v := infinite
	for _, i := range shuffledColumns() {
		if !s.IsValidMove(i) {
			continue
		}
		ss, _ := s.Move(i)
		if _, val := a.minmax(ss.NextTurn(), depth-1, !maxPlayer); val < v {
			v = val
			bestCol = i
		}
	}
	return bestCol, v
}

// alphabeta implements the Min-Max algorithm with Alpha-Beta pruning, allowing
// the AI to search deeper down the game tree. Returns the best move and value
// of that move.
func (a *AI) alphabeta(s *game.State, depth int, α, β value, maxPlayer bool) (game.Column, value) {
	if depth == 0 || s.IsGameOver() {
		return game.MaxColumn, a.stateValue(s, depth)
	}
	var bestCol game.Column
	if maxPlayer {
		v := -infinite
		for _, i := range shuffledColumns() {
			if !s.IsValidMove(i) {
				continue
			}
			ss, _ := s.Move(i)
			if _, val := a.alphabeta(ss.NextTurn(), depth-1, α, β, !maxPlayer); val > v {
				v = val
				bestCol = i
			}
			if v > α {
				α = v
			}
			if β <= α {
				// Prune.
				break
			}
		}
		return bestCol, v
	}
	v := infinite
	for _, i := range shuffledColumns() {
		if !s.IsValidMove(i) {
			continue
		}
		ss, _ := s.Move(i)
		if _, val := a.alphabeta(ss.NextTurn(), depth-1, α, β, !maxPlayer); val < v {
			v = val
			bestCol = i
		}
		if v < β {
			β = v
		}
		if β <= α {
			// Prune.
			break
		}
	}
	return bestCol, v
}

// shuffledColumns returns the shuffled game columns to introduce some
// randomness in the way the AI plays.
func shuffledColumns() []game.Column {
	var cols []game.Column
	for i := game.Column(0); i <= game.MaxColumn; i++ {
		cols = append(cols, i)
	}
	// Shuffle.
	for i := range cols {
		j := rand.Intn(i + 1)
		cols[i], cols[j] = cols[j], cols[i]
	}
	return cols
}