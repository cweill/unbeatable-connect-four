package game

import (
	"fmt"
)

type State struct {
	Grid [][]string
	Turn Player
}

type Player int

const (
        White Player = iota
        Black
)

func New() *State {
	return &State {
		Grid: func () [][]string {
			g := make([][]string, 6)
			for i := 0; i < len(g); i++ {
				g[i] = make([]string, 7)
			}
			return g
		}(),
	}
}

func (s *State) IsGameOver() bool    {
	return false
}

func (s *State) Move(i int) error {
	if i < 0 || i >= 7 {
		return fmt.Errorf("invalid move")
	}
	chip := "x"
	if s.Turn == Black {
		chip = "o"
	}
	for _, row := range s.Grid {
		if row[i] == "" {
			row[i] = chip
			return nil
		}
	}
	return fmt.Errorf("column is full")
}

func (s *State) String() string {
	res := ""
	for i := len(s.Grid)-1; i >= 0; i-- {
		row := s.Grid[i]
		res += "|"
		for _, v := range row {
			if v == "" {
				res += " |"
			} else {
				res += v + "|"
			}
		}
		res += "\n"
	}
	return res
}
