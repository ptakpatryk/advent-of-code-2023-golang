package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ptakpatryk/advent-of-code-2023-golang/lib"
)

type Pipes struct {
	grid           [][]string
	start          Point
	currentPositon Point
	prevPosition   Point
}

type Point [2]int

func (p *Pipes) move() {
	moveChar := p.currentChar()
	diff := p.getDirectionDiff()

	switch moveChar {
	case "|":
		p.goNextPipe(diff)
	case "-":
		p.goNextPipe(diff)
	case "L":
		if diff.compare(Point{1, 0}) {
			p.goNextPipe(Point{0, 1})
		} else {
			p.goNextPipe(Point{-1, 0})
		}
	case "J":
		if diff.compare(Point{0, -1}) {
			p.goNextPipe(Point{1, 0})
		} else {
			p.goNextPipe(Point{0, 1})
		}
	case "7":
		if diff.compare(Point{-1, 0}) {
			p.goNextPipe(Point{0, -1})
		} else {
			p.goNextPipe(Point{1, 0})
		}
	case "F":
		if diff.compare(Point{0, 1}) {
			p.goNextPipe(Point{-1, 0})
		} else {
			p.goNextPipe(Point{0, -1})
		}
	}
}

func (p *Pipes) startMove() {
	startChars := map[Point][]string{
		{-1, 0}: {"-", "F", "L"},
		{0, -1}:  {"|", "F", "7"},
		{1, 0}:  {"-", "J", "7"},
		{0, 1}: {"|", "J", "L"},
	}

	for startPoint, starts := range startChars {
		if startPoint[0] < 0 || startPoint[1] > len(p.grid[0])-1 || startPoint[1] < 0 || startPoint[1] > len(p.grid)-1 {
			continue
		}
		potentialNextPoint := Point{p.start[0] + startPoint[0], p.start[1] + startPoint[1]}
		potentialNextChar := p.grid[potentialNextPoint[1]][potentialNextPoint[0]]
		if slices.Index(starts, potentialNextChar) != -1 {
			p.currentPositon = potentialNextPoint
      fmt.Println(potentialNextChar)
			return
		}
	}
}

func (p *Pipes) currentChar() string {
	currX, currY := p.currentPositon[0], p.currentPositon[1]
	currChar := p.grid[currY][currX]

	return currChar
}

func (p *Point) compare(to Point) bool {
	return p[0] == to[0] && p[1] == to[1]
}

func (p *Pipes) goNextPipe(by Point) {
	t := p.currentPositon
	p.currentPositon = Point{p.currentPositon[0] - by[0], p.currentPositon[1] - by[1]}
	p.prevPosition = t
}

func (p *Pipes) getDirectionDiff() Point {
	currX, currY := p.currentPositon[0], p.currentPositon[1]
	prevX, prevY := p.prevPosition[0], p.prevPosition[1]

	return Point{prevX - currX, prevY - currY}
}

func main() {
	var input, err = os.ReadFile("day_10/input.txt")
	lib.CheckError(err)
	fmt.Println(partOne(string(input)))
	/* fmt.Println(partTwo(string(input))) */
}

func partOne(fileInput string) int {
	grid := getGrid(string(fileInput))
	grid.startMove()
	var moves = 1
	fmt.Printf("%+v \n", grid.currentPositon)

	for grid.currentChar() != "S" {
		grid.move()
		moves++
	}

	fmt.Println(moves)

	return moves / 2
}

func partTwo(fileInput string) int {
	return 0
}

func getGrid(input string) *Pipes {
	lines := strings.Split(input, "\n")
	var start Point
	grid := make([][]string, len(lines), len(lines))

	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		for j, ch := range l {
			if ch == 'S' {
				start[0] = j
				start[1] = i
			}
			grid[i] = append(grid[i], string(ch))
		}
	}

	return &Pipes{grid: grid, start: start, currentPositon: start, prevPosition: start}
}
