package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"golang.org/x/term"
)

func main() {
	cols, rows, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		cols = 20
		rows = 80
	}

	game := GameOfLife{Cols: cols / 2, Rows: rows - 1}
	game.Init()

	// Fill Board.
	rand.Seed(time.Now().UnixNano())

	for col := range game.board {
		for row := range game.board[col] {
			if randBool() {
				game.ToggleCell(col, row)
			}
		}
	}

	// Hide Cursor.
	fmt.Print("\033[?25l")

	// CTRL+C Signal.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	// Evolve Speed.
	ticker := time.NewTicker(250 * time.Millisecond)

Loop:
	for {
		select {
		case <-sig:
			// Show Cursor.
			fmt.Print("\033[?25h")

			break Loop
		case <-ticker.C:
			printBoard(game)
			game.Evolve()
		}
	}

}

func randBool() bool {
	return rand.Intn(7) == 1
}

func printBoard(g GameOfLife) {
	// Prevent flicker.
	fmt.Print("\033[H")

	for row := 0; row < g.Rows; row++ {
		for col := 0; col < g.Cols; col++ {
			if g.board[col][row] {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}

		fmt.Print("\n")
	}
}
