package main

import (
	"bufio"
	"fmt"
	"github.com/sanan-fataliyev/i2048/grid"
	"os"
)

func clear()  {
	fmt.Print("\033[H\033[2J")
}
func main() {
	g := grid.NewGrid(4, 4)

	// To create dynamic array
	scanner := bufio.NewScanner(os.Stdin)

	for {
		clear()
		g.Print()
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()

		switch text {
		case "q":
			return
		case "w":
			g.SwipeUp()
		case "a":
			g.SwipeLeft()
		case "d":
			g.SwipeRight()
		case "s":
			g.SwipeDown()
		}
		g.AddNewCells()
	}
}
