// Package game contains all Connect-Four state, player, and rendering logic.
package game

import (
	"errors"
)

// Player represents the current turn.
type Player int

const (
	// White is represented by an "X".
	White Player = iota
	// Black is represented by an "O".
	Black
)

// String returns the string representation of a player.
func (p Player) String() string {
	if p == White {
		return "X"
	}
	return "O"
}

// State represents the Connect-Four game's current state.
type State struct {
	// Grid represents the chips in the board.
	Grid [][]string
	// Turn is the current player's turn.
	Turn Player
}

// New returns a new Connect-Four game.
func New() *State {
	return &State{
		Grid: func() [][]string {
			g := make([][]string, 6)
			for i := 0; i < len(g); i++ {
				g[i] = make([]string, 7)
			}
			return g
		}(),
		Turn: White,
	}
}

// copy creates a deep copy of a given state.
func (s *State) copy() *State {
	var g [][]string
	for _, col := range s.Grid {
		var c []string
		for _, v := range col {
			c = append(c, v)
		}
		g = append(g, c)
	}
	return &State{
		Grid: g,
		Turn: s.Turn,
	}
}

// IsGameOver returns whether a player won or whether they reached a stalemate.
func (s *State) IsGameOver() bool {
	var freeSpace bool
	for i := 0; i < len(s.Grid); i++ {
		for j := 0; j < len(s.Grid[i]); j++ {
			v := s.Grid[i][j]
			if v == "" {
				freeSpace = true
				continue
			}
			if j < len(s.Grid[i])-3 && v == s.Grid[i][j+1] && v == s.Grid[i][j+2] && v == s.Grid[i][j+3] {
				return true
			}
			if i < len(s.Grid)-3 && v == s.Grid[i+1][j] && v == s.Grid[i+2][j] && v == s.Grid[i+3][j] {
				return true
			}
			if i < len(s.Grid)-3 && j < len(s.Grid[i])-3 && v == s.Grid[i+1][j+1] && v == s.Grid[i+2][j+2] && v == s.Grid[i+3][j+3] {
				return true
			}
			if i < len(s.Grid)-3 && j >= 3 && v == s.Grid[i+1][j-1] && v == s.Grid[i+2][j-2] && v == s.Grid[i+3][j-3] {
				return true
			}
		}
	}
	return !freeSpace
}

// Column is the selected column to drop the player's chip.
type Column int

// MaxColumn is index of the last column.
const MaxColumn = Column(6)

var (
	InvalidMoveError = errors.New("invalid move")
	ColumnFullError  = errors.New("column is full")
)

func (s *State) IsValidMove(c Column) bool {
	if c < 0 || c > MaxColumn {
		return false
	}
	for _, row := range s.Grid {
		if row[c] == "" {
			return true
		}
	}
	return false
}

// Move places a chip in the selected column. It does not mutate the original
// object, instead it copies the state and returns a new state with the move
// applied.
func (s *State) Move(c Column) (*State, error) {
	if c < 0 || c > MaxColumn {
		return nil, InvalidMoveError
	}
	cp := s.copy()
	for _, row := range cp.Grid {
		if row[c] == "" {
			row[c] = cp.Turn.String()
			return cp, nil
		}
	}
	return nil, ColumnFullError
}

// NextTurn returns a new state where it's the next player's turn. It does not
// mutate the original object, instead it copies the state and returns a new
// state with the next player up.
func (s *State) NextTurn() *State {
	cp := s.copy()
	if cp.Turn == White {
		cp.Turn = Black
	} else {
		cp.Turn = White
	}
	return cp
}

// String returns a string representation of the game board.
func (s *State) String() string {
	res := "\n"
	for i := len(s.Grid) - 1; i >= 0; i-- {
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
	res += "|1|2|3|4|5|6|7|\n"
	return res
}
