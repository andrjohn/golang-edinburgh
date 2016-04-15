package main

import (
	_ "fmt"
	tb "github.com/nsf/termbox-go"
	"math"
	"math/rand"
	"time"
)

func renderBoard(board [][]bool) {
	tb.Clear(tb.ColorWhite, tb.ColorBlack)
	for y, _ := range board {
		for x, v := range board[y] {
			if v {
				tb.SetCell(x, y, ' ', tb.ColorRed, tb.ColorRed)
			} else {
				tb.SetCell(x, y, ' ', tb.ColorBlack, tb.ColorBlack)
			}
		}
	}
	tb.Flush()
}

func newBoard(width, height int) [][]bool {
	board := make([][]bool, height, height)
	for y, _ := range board {
		row := make([]bool, width, width)
		for x, _ := range row {
			row[x] = (rand.Float32() > 0.5)
		}
		board[y] = row
	}
	return board
}

func getIndex(index, modifier, max int) int {
	//return (index + modifier) % max
	return int(math.Mod(float64(index+modifier+max), float64(max)))
}

func countNeighbours(board [][]bool, x, y int) int {
	var count int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			new_y := getIndex(y, i, len(board))
			new_x := getIndex(x, j, len(board[y]))
			if i == 0 && j == 0 {
				continue
			}
			if board[new_y][new_x] {
				count++
			}
		}
	}
	return count
}

func calcBoard(in, out [][]bool) {
	for y, _ := range in {
		for x, alive := range in[y] {
			neighbours := countNeighbours(in, x, y)
			if alive {
				switch neighbours {
				case 2, 3:
					out[y][x] = true
				default:
					out[y][x] = false
				}
			} else if neighbours == 3 {
				out[y][x] = true
			} else {
				out[y][x] = false
			}
		}
	}
}

func main() {
	if err := tb.Init(); err != nil {
		panic(err)
	}
	defer tb.Close()

	board1 := newBoard(tb.Size())
	board2 := newBoard(tb.Size())

	for {
		renderBoard(board1)
		time.Sleep(500 * time.Millisecond)
		calcBoard(board1, board2)
		board1, board2 = board2, board1
	}
}
