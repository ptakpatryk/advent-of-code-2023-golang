package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
  "github.com/ptakpatryk/advent-of-code-2023-golang/lib"
)

var input, err = os.ReadFile("day_01/input.txt")

func main() {
  lib.CheckError(err)
  fmt.Println(partOne(string(input)))
}

func partOne(fileInput string) int {
  lines := strings.Split(fileInput, "\n")

  var numOuters [][]string

  for _, l := range lines {
    var lineNumbers []string
    for _, c := range l {
      letter := string(c)
      _, err := strconv.Atoi(letter)
      if err != nil {
        continue;
      }

      lineNumbers = append(lineNumbers, letter)
    }

    if len(lineNumbers) > 0 {
      numOuters = append(numOuters, []string{lineNumbers[0], lineNumbers[len(lineNumbers) - 1]})
    }
  }

  var sum int

  for _, numSet := range numOuters {
    pairSum, err := strconv.Atoi(numSet[0] + numSet[1])
    if err != nil {
      continue
    }
    sum += pairSum
  }

  return sum
}

