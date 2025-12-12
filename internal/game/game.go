package game

import "fmt"

//player marks
const (
	Empty = 0
	X     = 1
	O     = 2
)

type Board [9]int

func NewBoard() Board { return Board{} }

func (b *Board) Reset() { *b = NewBoard() }

func (b Board) Get(r, c int) int {
	return b[r*3+c]
}

func (b *Board) Set(r, c, val int) error {

	if r < 0 || r > 2 || c < 0 || c > 2 {
		return fmt.Errorf("index out of range")
	}

	if b[r*3+c] != Empty {
		return fmt.Errorf("cell occupied")
	}
	b[r*3+c] = val
	return nil
}

func (b Board) Winner() (int, bool) {
	wins := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // cols
		{0, 4, 8}, {2, 4, 6}, // diags
	}
	for _, w := range wins {
		a, bb, c := b[w[0]], b[w[1]], b[w[2]]
		if a != Empty && a == bb && a == c {
			return a, true
		}
	}
	return Empty, false
}

func (b Board) Full() bool {
	for _, v := range b {
		if v == Empty {
			return false
		}
	}
	return true
}

// Moves function returns available moves
func (b Board) Moves() []int {
	out := []int{}
	for i, v := range b {
		if v == Empty {
			out = append(out, i)
		}
	}
	return out
}

// MakeMoves function marks player moves on the board
func (b *Board) MakeMove(idx int, player int) error {
	if idx < 0 || idx > 8 {
		return fmt.Errorf("index out of range")
	}

	if b[idx] != Empty {
		return fmt.Errorf("cell occupied")
	}
	b[idx] = player
	return nil
}

// Pretty returns a simple ascii view used for debug or logs
func (b Board) Pretty() string {
	s := ""
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			v := b.Get(r, c)
			ch := "."
			switch v {
			case X:
				ch = "X"
			case O:
				ch = "O"
			}
			s += ch
			if c < 2 {
				s += " "
			}
		}
		if r < 2 {
			s += "\n"
		}
	}
	return s
}
