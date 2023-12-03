package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ptakpatryk/advent-of-code-2023-golang/lib"
)

var input, err = os.ReadFile("day_03/input.txt")

func main() {
	lib.CheckError(err)
	fmt.Println(partOne(string(input)))
	fmt.Println(partTwo(string(input)))
}

func partOne(fileInput string) int {
	matrix := getMatrix(fileInput)
	var eligibleNumbers []string

	for y, line := range matrix {
		if len(line) == 0 {
			continue
		}

		var currentNumber []string
		var isCurrAdjacent bool

		for x, char := range line {
			_, err := strconv.Atoi(char)
			if err != nil {
				if len(currentNumber) != 0 && isCurrAdjacent {
					eligibleNumbers = append(eligibleNumbers, strings.Join(currentNumber, ""))
				}
				isCurrAdjacent = false
				currentNumber = []string{}
				continue
			}
			currentNumber = append(currentNumber, char)

			if isAdjacentTo(x, y, matrix, isSymbol) {
				isCurrAdjacent = true
			}
		}

		if len(currentNumber) != 0 && isCurrAdjacent {
			eligibleNumbers = append(eligibleNumbers, strings.Join(currentNumber, ""))
		}
	}

	var sum int
	for _, n := range eligibleNumbers {
		num, _ := strconv.Atoi(n)
		sum += num
	}
	return sum
}

func partTwo(fileInput string) int {
	matrix := getMatrix(fileInput)
	var possibleGears = make(map[string][]string)

	for y, line := range matrix {
		if len(line) == 0 {
			continue
		}

		var currentNumber []string
		var isCurrAdjacent bool
		var adjacentStars = make(map[string]bool)

		for x, char := range line {
			_, err := strconv.Atoi(char)
			if err != nil {
				if len(currentNumber) != 0 && isCurrAdjacent {
					for starPoint := range adjacentStars {
						possibleGears[starPoint] = append(possibleGears[starPoint], strings.Join(currentNumber, ""))
					}
				}
				adjacentStars = make(map[string]bool)
				isCurrAdjacent = false
				currentNumber = []string{}
				continue
			}
			currentNumber = append(currentNumber, char)

			if isAdjacentTo(x, y, matrix, isSymbol) {
				isCurrAdjacent = true
			}

			for _, starPoint := range getAdjacentStars(x, y, matrix) {
				adjacentStars[starPoint] = true
			}
		}

		if len(currentNumber) != 0 && isCurrAdjacent {
			for starPoint := range adjacentStars {
				possibleGears[starPoint] = append(possibleGears[starPoint], strings.Join(currentNumber, ""))
			}
		}
	}

	var sum int

	for _, values := range possibleGears {
		if len(values) == 2 {
			v1, _ := strconv.Atoi(values[0])
			v2, _ := strconv.Atoi(values[1])
			sum = sum + (v1 * v2)
		}
	}

	return sum
}

func isAdjacentTo(x int, y int, matrix [][]string, matcher func(string) bool) bool {
	adjacentPoints := getAdjacentPoints(x, y)

	for i, pointsSet := range adjacentPoints {
		xn, yn := pointsSet[0], pointsSet[1]
		if xn < 0 || yn < 0 || yn > len(matrix)-1 || xn > len(matrix[i])-1 {
			continue
		}
		v := matrix[yn][xn]
		if matcher(v) {
			return true
		}
	}

	return false
}

func getAdjacentStars(x int, y int, matrix [][]string) []string {
	var starPoints []string
	adjacentPoints := getAdjacentPoints(x, y)

	for i, pointsSet := range adjacentPoints {
		xn, yn := pointsSet[0], pointsSet[1]
		if xn < 0 || yn < 0 || yn > len(matrix)-1 || xn > len(matrix[i])-1 {
			continue
		}
		v := matrix[yn][xn]
		if isStar(v) {
			ys, xs := strconv.Itoa(yn), strconv.Itoa(xn)
			starPoints = append(starPoints, strings.Join([]string{xs, ys}, ","))
		}
	}

	return starPoints
}

func getAdjacentPoints(x int, y int) [][]int {
	return [][]int{
		{x, y - 1},
		{x + 1, y - 1},
		{x + 1, y},
		{x + 1, y + 1},
		{x, y + 1},
		{x - 1, y + 1},
		{x - 1, y},
		{x - 1, y - 1},
	}
}

func getMatrix(input string) [][]string {
	var matrix [][]string
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		matrix = append(matrix, []string{})
		for _, char := range line {
			matrix[i] = append(matrix[i], string(char))
		}
	}

	return matrix
}

func isStar(s string) bool {
	return s == "*"
}

func isSymbol(s string) bool {
	return !isNumberStr(s) && s != "."
}

func isNumberStr(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return true
}
