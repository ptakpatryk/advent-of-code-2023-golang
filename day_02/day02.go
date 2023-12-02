package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ptakpatryk/advent-of-code-2023-golang/lib"
)

var input, err = os.ReadFile("day_02/input.txt")
var cubes = [3]string{"blue", "green", "red"}

type Game struct {
	Green int
	Red   int
	Blue  int
}

func (g Game) isWithinLimit() bool {
  return !(g.Red > 12 || g.Blue > 14 || g.Green > 13)
}

func (g *Game) mergeIfBigger(round Game) {
	g.Green = max(g.Green, round.Green)
	g.Red = max(g.Red, round.Red)
	g.Blue = max(g.Blue, round.Blue)
}

func main() {
	lib.CheckError(err)
	fmt.Println(partOne(string(input)))
	fmt.Println(partTwo(string(input)))
}

func partOne(fileInput string) int {
	lines := strings.Split(fileInput, "\n")
	var games = make(map[string]*Game)

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		var id = strconv.Itoa(i + 1)
		games[id] = &Game{}

		game := strings.Split(line, ":")
		game = strings.Split(game[1], ";")

		rounds := getRounds(game)

		for _, round := range rounds {
			if !round.isWithinLimit() {
				delete(games, id)
				break
			}
		}
	}

	var sum int
	for k := range games {
		id, _ := strconv.Atoi(k)
		sum += id
	}
	return sum
}

func partTwo(fileInput string) int {
	lines := strings.Split(fileInput, "\n")
	var games = make(map[string]*Game)

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		var id = strconv.Itoa(i + 1)
		games[id] = &Game{}

		game := strings.Split(line, ":")
		game = strings.Split(game[1], ";")

		rounds := getRounds(game)

		for _, round := range rounds {
			games[id].mergeIfBigger(round)
		}
	}

	var sum int
	for id := range games {
		power := games[id].Red * games[id].Blue * games[id].Green
		sum += power
	}
	return sum
}

func getRounds(game []string) []Game {
	var rounds []Game
	for _, gameCubes := range game {
		cubes := strings.Split(gameCubes, ",")
		var roundResult = make(map[string]int)

		for _, cube := range cubes {
			cube = strings.Trim(cube, " ")
			cubeSplit := strings.Split(cube, " ")

			color := cubeSplit[1]
			val, err := strconv.Atoi(cubeSplit[0])
			lib.CheckError(err)

			roundResult[color] = val
		}

		jsonGame, err := json.Marshal(roundResult)
		lib.CheckError(err)

		var round = Game{}
		if err := json.Unmarshal(jsonGame, &round); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		rounds = append(rounds, round)
	}

	return rounds
}
