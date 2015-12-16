package game

import (
	"time"
)

type AI struct {}

func (a *AI) ChooseMove(s *State) Column {
	time.Sleep(500 * time.Millisecond)
	return 1
} 