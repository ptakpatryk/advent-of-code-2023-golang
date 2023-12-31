package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ptakpatryk/advent-of-code-2023-golang/lib"
)

var input, err = os.ReadFile("day_06/input.txt")

func main() {
	lib.CheckError(err)
	fmt.Println(partOne(string(input)))
	fmt.Println(partTwo(string(input)))
}

type Race struct {
	time     int
	distance int
}

func partOne(fileInput string) int {
	var races []Race

	times := strings.Fields(strings.Split(fileInput, "\n")[0])
	distances := strings.Fields(strings.Split(fileInput, "\n")[1])

	for i, time := range times {
		t, err := strconv.Atoi(time)
		if err != nil {
			continue
		}
		d, _ := strconv.Atoi(distances[i])

		races = append(races, Race{t, d})
	}

	return getWaysToWin(races)
}

func partTwo(fileInput string) int {
	time := strings.Split(strings.Split(fileInput, "\n")[0], ":")[1]
	time = strings.ReplaceAll(time, " ", "")
	distance := strings.Split(strings.Split(fileInput, "\n")[1], ":")[1]
	distance = strings.ReplaceAll(distance, " ", "")

	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(distance)

	return getWaysToWin([]Race{{t, d}})
}

func getWaysToWin(races []Race) int {
	var waysToWin []int

	for _, race := range races {
		var waysToWinRace int
		for i := 1; i < race.time; i++ {
			if willBeat(i, race) {
				waysToWinRace++
			}
		}

		waysToWin = append(waysToWin, waysToWinRace)
	}

	result := 1
	for _, ways := range waysToWin {
		result *= ways
	}

	return result
}

func willBeat(speed int, race Race) bool {
	distance := speed * (race.time - speed)

	return distance > race.distance
}
