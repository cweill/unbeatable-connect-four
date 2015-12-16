package ai

import (
	"c4/game"
	"math/rand"
	"fmt"
)

const maxDepth = 5

type AI struct {
	Player game.Player
}

func (a *AI) ChooseMove(s *game.State) game.Column {
	fmt.Println(s.Turn)
	col, _ := a.minmax(s, maxDepth, true)
	return col
}

type value int

const (
	infinite = value(99999999)
	uncertain = value(0)
	win = value(10)
	loss = value(-10)
)

func (a *AI) stateValue(s *game.State, depth int) value {
	if !s.IsGameOver() {
		return uncertain
	}
	v := value(depth + 1)
	if s.Turn == a.Player {
		return win * v
	}
	return loss * v
}

func (a *AI) minmax(s *game.State, depth int, maxPlayer bool) (game.Column, value) {
	if depth == 0 || s.IsGameOver() {
		return game.Column(rand.Intn(int(game.MaxColumn))), a.stateValue(s, depth)
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
