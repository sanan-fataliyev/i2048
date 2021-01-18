package grid

import (
	"fmt"
	"math/rand"
)

//type direction struct {
//	DeltaX, DeltaY int
//}
//
//var (
//	UP    = direction{DeltaX: 0, DeltaY: 1}
//	RIGHT = direction{DeltaX: 1, DeltaY: 0}
//	DOWN  = direction{DeltaX: 0, DeltaY: -1}
//	LEFT  = direction{DeltaX: -1, DeltaY: 0}
//)

type Grid struct {
	RowSize int
	ColSize int
	Cells   [][]int
}

func NewGrid(rows, cols int) *Grid {
	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]int, cols)
	}
	return &Grid{
		RowSize: rows,
		ColSize: cols,
		Cells:   grid,
	}
}

func FromMatrix(m [][]int) *Grid {
	return &Grid{
		RowSize: len(m),
		ColSize: len(m[0]),
		Cells:   m,
	}
}

func (g *Grid) SwipeUp() {

	for r := 0; r < g.RowSize-1; r++ {
		for c := 0; c < g.ColSize; c++ {
			// if the cell is empty
			if g.Cells[r][c] == 0 {
				downR := g.nextDown(r, c)
				if downR == -1 {
					continue // skip this column
				} else {
					// swap them
					g.Cells[r][c], g.Cells[downR][c] = g.Cells[downR][c], g.Cells[r][c]
				}
			}
			curr := g.Cells[r][c]
			downR := g.nextDown(r, c)
			if downR == -1 {
				continue // skip this column
			}
			down := g.Cells[downR][c]

			// merge
			if down == curr {
				g.Cells[r][c] *= 2    // double
				g.Cells[downR][c] = 0 // flush
			}
		}
	}
}

func (g *Grid) SwipeRight() {
	for c := g.ColSize - 1; c > 0; c-- {
		for r := 0; r < g.RowSize; r++ {
			// if the cell is empty
			if g.Cells[r][c] == 0 {
				// try to find the next nonempty cell
				leftC := g.nextLeft(r, c)
				if leftC == -1 {
					continue // skip this row
				} else {
					// swap them
					g.Cells[r][c], g.Cells[r][leftC] = g.Cells[r][leftC], g.Cells[r][c]
				}
			}
			curr := g.Cells[r][c]
			leftC := g.nextLeft(r, c)
			if leftC == -1 {
				continue // skip this column
			}
			left := g.Cells[r][leftC]

			// merge
			if left == curr {
				g.Cells[r][c] *= 2    // double
				g.Cells[r][leftC] = 0 // flush
			}
		}
	}
}

/*
2 0 0 0
2 0 0 0
2 0 0 0
2 0 0 0
=>
0 0 0 0
0 0 0 0
4 0 0 0
4 0 0 0
*/

func (g Grid) nextUp(r0, c0 int) int {
	upR, _ := g.nextNonZero(r0, c0, -1, 0)
	return upR
}

func (g Grid) nextRight(r0, c0 int) int {
	_, c := g.nextNonZero(r0, c0, 0, 1)
	return c
}

func (g Grid) nextDown(r0, c0 int) int {
	r, _ := g.nextNonZero(r0, c0, 1, 0)
	return r
}

func (g Grid) nextLeft(r0, c0 int) int {
	_, c := g.nextNonZero(r0, c0, 0, -1)
	return c
}

// generic finder
func (g Grid) nextNonZero(r0, c0, deltaR, deltaC int) (r, c int) {
	for ri, ci := r0+deltaR, c0+deltaC; ri >= 0 && ri < g.RowSize && ci >= 0 && ci < g.ColSize; {
		if g.Cells[ri][ci] != 0 {
			return ri, ci
		}
		ri += deltaR
		ci += deltaC
	}
	return -1, -1
}

// done
func (g *Grid) SwipeDown() {

	for r := g.RowSize - 1; r > 0; r-- {
		for c := 0; c < g.ColSize; c++ {
			// if the cell is empty
			if g.Cells[r][c] == 0 {
				// try to find the next nonempty cell on top
				upR := g.nextUp(r, c)
				if upR == -1 {
					continue // skip this column
				} else {
					// swap them
					g.Cells[r][c], g.Cells[upR][c] = g.Cells[upR][c], g.Cells[r][c]
				}
			}
			curr := g.Cells[r][c]
			upR := g.nextUp(r, c)
			if upR == -1 {
				continue // skip this column
			}
			up := g.Cells[upR][c]

			// merge
			if up == curr {
				g.Cells[r][c] *= 2  // double
				g.Cells[upR][c] = 0 // flush
			}
		}
	}
}

func (g Grid) SwipeLeft() {
	for c := 0; c < g.ColSize-1; c++ {
		for r := 0; r < g.RowSize; r++ {
			// if the cell is empty
			if g.Cells[r][c] == 0 {
				// try to find the next nonempty cell
				rightC := g.nextRight(r, c)
				if rightC == -1 {
					continue // skip this row
				} else {
					// swap them
					g.Cells[r][c], g.Cells[r][rightC] = g.Cells[r][rightC], g.Cells[r][c]
				}
			}
			curr := g.Cells[r][c]
			rightC := g.nextRight(r, c)
			if rightC == -1 {
				continue // skip this column
			}
			right := g.Cells[r][rightC]

			// merge
			if right == curr {
				g.Cells[r][c] *= 2     // double
				g.Cells[r][rightC] = 0 // flush
			}
		}
	}

}

func (g Grid) Print() {

	for _, row := range g.Cells {
		for _, cell := range row {
			fmt.Printf("|%3d ", cell)
		}
		fmt.Println("|")
	}
}

func (g Grid) IsGameOver() bool {
	return false
}

func (g *Grid) AddNewCells()  {
	r, c := rand.Intn(g.RowSize), rand.Intn(g.ColSize);
	for  g.Cells[r][c] != 0 {
		r, c = rand.Intn(g.RowSize), rand.Intn(g.ColSize)
	}
	g.Cells[r][c] = 2
}