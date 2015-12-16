package game

import (
	"fmt"
)

type State struct {
	Grid [][]string
	Turn Player
}

type Player int

func (p Player) String() string {
	if p == White {
		return "X"
	}
	return "O"
}

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
		Turn: White,
	}
}

func (s *State) IsGameOver() bool    {
	var freeSpace bool
	for i := 0; i < len(s.Grid); i++ {
		for j := 0; j < len(s.Grid[i]); j++ {
			v := s.Grid[i][j]
			if v == "" {
				freeSpace = true
				continue
			}
			if j < len(s.Grid[i]) - 3 && v == s.Grid[i][j+1] && v == s.Grid[i][j+2] && v == s.Grid[i][j+3] {
				return true
			}
			if i < len(s.Grid) - 3 &&v == s.Grid[i+1][j] && v == s.Grid[i+2][j] && v == s.Grid[i+3][j] {
				return true
			}
			if i < len(s.Grid) - 3 && j < len(s.Grid[i]) - 3 && v == s.Grid[i][j+1] && v == s.Grid[i+2][j+2] && v == s.Grid[i+3][j+3] {
				return true
			}
		}
	}
	return !freeSpace
}

func (s *State) Move(i int) error {
	if i < 0 || i >= 7 {
		return fmt.Errorf("invalid move")
	}
	for _, row := range s.Grid {
		if row[i] == "" {
			row[i] = s.Turn.String()
			return nil
		}
	}
	return fmt.Errorf("column is full")
}

func (s *State) NextTurn() {
	if s.Turn == White {
		s.Turn = Black
	} else {
		s.Turn = White
	}
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
