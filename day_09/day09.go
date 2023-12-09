package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ptakpatryk/advent-of-code-2023-golang/lib"
)

var input, err = os.ReadFile("day_09/input.txt")

func main() {
	lib.CheckError(err)
	fmt.Println(partOne(string(input)))
	fmt.Println(partTwo(string(input)))
}

func partOne(fileInput string) int {
	lines := strings.Split(fileInput, "\n")
	var sum int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var nums []int

		numsStr := strings.Fields(line)
		for _, n := range numsStr {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		var sequences [][]int
		sequences = append(sequences, nums)
		for i := 0; !isEnd(sequences[i]); i++ {
			sequences = append(sequences, getNextSequence(sequences[i]))
		}

		sum += getHistoryValue(sequences)
	}

	return sum
}

func partTwo(fileInput string) int {
	lines := strings.Split(fileInput, "\n")
	var sum int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var nums []int

		numsStr := strings.Fields(line)
		for _, n := range numsStr {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		var sequences [][]int
		sequences = append(sequences, nums)
		for i := 0; !isEnd(sequences[i]); i++ {
			sequences = append(sequences, getNextSequence(sequences[i]))
		}

		sum += getHistoryValueReversed(sequences)
	}

	return sum
}

func getHistoryValue(sequences [][]int) int {
	sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)
	for i := len(sequences) - 2; i >= 0; i-- {
		lastNum := sequences[i][len(sequences[i])-1] + sequences[i+1][len(sequences[i])-1]
		sequences[i] = append(sequences[i], lastNum)
	}

	return sequences[0][len(sequences[0])-1]
}

func getHistoryValueReversed(sequences [][]int) int {
	sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)
	for i := len(sequences) - 2; i >= 0; i-- {
		firstNum := sequences[i][0] - sequences[i+1][0]
		sequences[i] = append([]int{firstNum}, sequences[i]...)
	}

	return sequences[0][0]
}

func getNextSequence(numbers []int) []int {
	var sequence []int
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		sequence = append(sequence, diff)
	}

	return sequence
}

func isEnd(numbers []int) bool {
	for _, n := range numbers {
		if n != 0 {
			return false
		}
	}
	return true
}
