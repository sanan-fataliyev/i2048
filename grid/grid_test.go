package grid

import "testing"

func gridEQ(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	for ri := range a {
		if len(a[ri]) != len(b[ri]) {
			return false
		}
		for ci := range a[ri] {
			if a[ri][ci] != b[ri][ci] {
				return false
			}
		}
	}

	return true
}

func TestGrid_SwipeDown(t *testing.T) {

	grid := FromMatrix([][]int{
		{2, 2, 2, 2, 2, 0, 2},
		{2, 0, 0, 0, 2, 0, 4},
		{2, 0, 0, 0, 2, 0, 0},
		{2, 0, 2, 0, 0, 0, 2},
	})

	grid.SwipeDown()

	expected := FromMatrix([][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 2},
		{4, 0, 0, 2, 0, 4},
		{4, 2, 4, 4, 0, 2},
	})

	if !gridEQ(grid.Cells, expected.Cells) {
		println("expected:")
		expected.Print()
		println("actual:")
		grid.Print()
		t.Fail()
	}

}

func TestGrid_SwipeUp(t *testing.T) {

	grid := FromMatrix([][]int{
		{2, 0, 2, 0, 0, 2},
		{2, 0, 0, 2, 0, 0},
		{2, 0, 0, 2, 0, 4},
		{2, 2, 2, 2, 0, 2},
	})

	grid.Print()
	grid.SwipeUp()

	expected := FromMatrix([][]int{
		{4, 2, 4, 2, 4, 0, 2},
		{4, 0, 0, 0, 2, 0, 4},
		{0, 0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 0, 0, 0},
	})

	if !gridEQ(grid.Cells, expected.Cells) {
		println("expected:")
		expected.Print()
		println("actual:")
		grid.Print()
		t.Fail()
	}

}

func TestGrid_SwipeRight(t *testing.T) {

	grid := FromMatrix([][]int{
		{2, 2, 2, 2},
		{2, 0, 0, 0},
		{2, 0, 0, 2},
		{2, 2, 2, 0},
		{0, 0, 0, 0},
		{2, 4, 0, 2},
	})

	grid.Print()
	grid.SwipeRight()

	expected := FromMatrix([][]int{
		{0, 0, 4, 4},
		{0, 0, 0, 2},
		{0, 0, 0, 4},
		{0, 0, 2, 4},
		{0, 0, 0, 0},
		{0, 2, 4, 2},
	})

	if !gridEQ(grid.Cells, expected.Cells) {
		println("expected:")
		expected.Print()
		println("actual:")
		grid.Print()
		t.Fail()
	}

}

func TestGrid_SwipeLeft(t *testing.T) {

	grid := FromMatrix([][]int{
		{2, 2, 2, 2},
		{0, 0, 0, 2},
		{2, 0, 0, 2},
		{0, 2, 2, 2},
		{0, 0, 0, 0},
		{2, 0, 4, 2},
	})

	grid.Print()
	grid.SwipeLeft()

	expected := FromMatrix([][]int{
		{4, 4, 0, 0},
		{2, 0, 0, 0},
		{4, 0, 0, 0},
		{4, 2, 0, 0},
		{0, 0, 0, 0},
		{2, 4, 2, 0},
	})

	if !gridEQ(grid.Cells, expected.Cells) {
		println("expected:")
		expected.Print()
		println("actual:")
		grid.Print()
		t.Fail()
	}

}
