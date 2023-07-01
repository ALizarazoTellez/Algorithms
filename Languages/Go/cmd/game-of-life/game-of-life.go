package main

type GameOfLife struct {
	Cols, Rows int

	board [][]bool
}

func (g *GameOfLife) Init() {
	g.board = make([][]bool, g.Cols)
	for col := range g.board {
		g.board[col] = make([]bool, g.Rows)
	}
}

func (g *GameOfLife) ToggleCell(x, y int) (isLive bool) {
	if g.board[x][y] {
		g.board[x][y] = false
	} else {
		g.board[x][y] = true
	}

	return g.board[x][y]
}

func (g *GameOfLife) Evolve() {
	newBoard := make([][]bool, g.Cols)
	for col := range newBoard {
		newBoard[col] = make([]bool, g.Rows)
		copy(newBoard[col], g.board[col])
	}

	for col := range g.board {
		for row := range g.board[col] {
			population := g.getPopulation(col, row)

			if g.board[col][row] { // Is Live.
				if population > 3 || population <= 1 {
					//g.ToggleCell(col, row)
					newBoard[col][row] = false
				}
			} else { // Is Death.
				if population == 3 {
					//g.ToggleCell(col, row)
					newBoard[col][row] = true
				}
			}
		}
	}

	g.board = newBoard
}

func (g *GameOfLife) getPopulation(x, y int) int {
	// a b c
	// d   e
	// f g h
	var neighbors = [...][2]int{
		[2]int{-1, -1}, // a
		[2]int{0, -1},  // b
		[2]int{1, -1},  // c
		[2]int{-1, 0},  // d
		[2]int{1, 0},   // e
		[2]int{-1, 1},  // f
		[2]int{0, 1},   // g
		[2]int{1, 1},   // h
	}

	var population int

	for _, neighbor := range neighbors {
		nx, ny := neighbor[0]+x, neighbor[1]+y

		if nx < 0 || nx >= g.Cols || ny < 0 || ny >= g.Rows {
			continue
		}

		if g.board[nx][ny] {
			population++
		}
	}

	return population
}
