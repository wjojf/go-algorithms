package main	 

import (
	"errors"
	"fmt"
)

type sudokuBoard struct {
	Board [9][9]int
}

func (s sudokuBoard) nextEmpty() (int, int, error) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.Board[i][j] == 0 {
				return i, j, nil
			}
		}
	}

	return 0,0, errors.New("No empty cells left")
}

func (s sudokuBoard) isValidChoice(i,j,v int) bool {
	// check row 
	for _, x := range(s.Board[i]) {
		if x == v {
			return false 
		}
	}

	// check col
	for _, r := range(s.Board) {
		if r[j] == v {
			return false
		}
	}

	// check square
	_i := int(i/3)
	_j := int(j/3)

	for x := 3 * _i; x < 3 * (_i + 1); x++{
		for y := 3 * _j; y < 3 * (_j + 1) ; y++{
			if s.Board[x][y] == v {
				return false
			}
		}
	}

	return true
}

func (s sudokuBoard) solveBoard() {
	i, j, e := s.nextEmpty()
	if e != nil {
		return
	}

	var backtracing func(int, int) bool
	backtracing = func(i,j int) bool {
		for k := 1; k <= 9; k++ {
			if s.isValidChoice(i,j,k) {
				s.Board[i][j] = k
				fmt.Println(i, j, k)
				i, j, e := s.nextEmpty()
				if e != nil {
					return true
				}

				status := backtracing(i,j)
				if status {
					return true
				}

			}

			s.Board[i][j] = 0
		}

		return false
	}

	backtracing(i, j)
}

func (s sudokuBoard) printBoard() {
	for _, r := range(s.Board) {
		for _, x := range(r) {
			fmt.Printf("%v ", x)
		}
		fmt.Print("\n")
	}

	fmt.Println("------------------")
}

func main() {

	testBoard := [9][9]int{
		{6, 0, 0, 0, 0, 2, 5, 0, 0},
		{0, 1, 7, 5, 0, 0 ,0, 0, 0},
		{4, 0, 0, 0, 0, 0, 0, 2, 0},
		{0, 7, 0, 0, 2, 3, 0, 6, 0},
		{0, 0, 0, 0, 1, 0, 3, 0 ,0},
		{0, 0, 2, 0, 0, 5, 7, 0 ,0},
		{0, 0, 0, 4, 0, 0, 0, 0, 0},
		{0, 9, 5, 0, 0 ,0 ,0, 3, 0},
		{1, 0, 8, 0, 0, 0, 9, 0, 0},
	}

	s := sudokuBoard{testBoard}
	fmt.Println(s.isValidChoice(0, 3, 1))
	fmt.Println(s.isValidChoice(1, 4, 1))
	s.printBoard()
}