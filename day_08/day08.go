package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ptakpatryk/advent-of-code-2023-golang/lib"
)

var input, err = os.ReadFile("day_08/input.txt")

type Node struct {
	L string
	R string
}

func main() {
	lib.CheckError(err)
	fmt.Println(partOne(string(input)))
	fmt.Println(partTwo(string(input)))
}

func partOne(fileInput string) int {
	file := strings.Split(fileInput, "\n\n")
	instruction := file[0]
	nodesStr := strings.Split(file[1], "\n")
	nodes := make(map[string]Node)

	for _, nodeLine := range nodesStr {
		if len(nodeLine) == 0 {
			continue
		}
		fields := strings.Fields(nodeLine)
		l := fields[2]
		r := fields[3]
		nodes[fields[0]] = Node{l[1:4], r[:len(r)-1]}
	}

	return findRecursive("AAA", nodes, []string{"ZZZ"} , instruction, 0)
}

func partTwo(fileInput string) int {
	file := strings.Split(fileInput, "\n\n")
	instruction := file[0]
	nodesStr := strings.Split(file[1], "\n")
	nodes := make(map[string]Node)

	var starts []string
	var ends []string

	for _, nodeLine := range nodesStr {
		if len(nodeLine) == 0 {
			continue
		}
		fields := strings.Fields(nodeLine)
		if fields[0][2] == 'A' {
			starts = append(starts, fields[0])
		}
		if fields[0][2] == 'Z' {
			ends = append(ends, fields[0])
		}
		l := fields[2]
		r := fields[3]
		nodes[fields[0]] = Node{l[1:4], r[:len(r)-1]}
	}

	var pathsToZ []int

	for _, start := range starts {
		pathsToZ = append(pathsToZ, findRecursive(start, nodes, ends, instruction, 0))
	}

	return lcmm(pathsToZ)
}

func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	return a
}

func lcm(a, b int) int {
	return (a * b / gcd(a, b))
}

func lcmm(args []int) int {
  if len(args) == 2 {
    return lcm(args[0], args[1])
  } else {
    var arg0 = args[0]
    return lcm(arg0, lcmm(args[1:]))
  }
}

func findRecursive(start string, nodes map[string]Node, ends []string, instruction string, moveCount int) int {
	if slices.Contains(ends, start) {
		return moveCount
	}
	dir := instruction[moveCount%len(instruction)]
	if dir == 'L' {
		start = nodes[start].L
	} else {
		start = nodes[start].R
	}

	return findRecursive(start, nodes, ends, instruction, moveCount+1)
}
