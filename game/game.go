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

func (s *State) Copy() *State {
	return &State {
		Grid: s.Grid,
		Turn: s.Turn,
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

func (s *State) Move(i int) (*State, error) {
	if i < 0 || i >= 7 {
		return nil, fmt.Errorf("invalid move")
	}
	cp := s.Copy()
	for _, row := range cp.Grid {
		if row[i] == "" {
			row[i] = cp.Turn.String()
			return cp, nil
		}
	}
	return nil, fmt.Errorf("column is full")
}

func (s *State) NextTurn() *State {
	cp := s
	if cp.Turn == White {
		cp.Turn = Black
	} else {
		cp.Turn = White
	}
	return cp
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
