package main

import "fmt"

func main() {
	boardSize := 5

	// Read board
	board := make(map[int]map[int]bool)
	for y := 0; y < boardSize; y++ {
		board[y] = make(map[int]bool)
		var line string
		fmt.Scanln(&line)
		for x, r := range line {
			if r == 'k' {
				board[y][x] = true
			}
		}
	}

	type vec struct { x,y int }
	attackVectors := []vec{
		{1,2},
		{2,1},
		{2,-1},
		{1,-2},
		{-1,-2},
		{-2,-1},
		{-2,1},
		{-1,2},
	}
	n := 0
	for y, xs := range board {
		for x := range xs {
			n++
			for _, v := range attackVectors {
				xa := v.x+x
				ya := v.y+y
				if board[ya][xa] {
					fmt.Println("invalid")
					return
				}
			}
		}
	}
	if n == 9 {
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
	}
}
