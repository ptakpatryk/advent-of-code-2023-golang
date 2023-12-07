package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/ptakpatryk/advent-of-code-2023-golang/lib"
)

var input, err = os.ReadFile("day_07/input.txt")

func main() {
	lib.CheckError(err)
	fmt.Println(partOne(string(input)))
	fmt.Println(partTwo(string(input)))
}

var ranks = [7][]int{
	{1, 1, 1, 1, 1},
	{1, 1, 1, 2},
	{1, 2, 2},
	{1, 1, 3},
	{2, 3},
	{1, 4},
	{5},
}

type Hand struct {
	cards string
	bid   int
	rank  int
}

func partOne(fileInput string) int {
	lines := strings.Split(fileInput, "\n")

	var cardsByStrength = make(map[int][]Hand)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		cards := strings.Fields(line)[0]
		bid, _ := strconv.Atoi(strings.Fields(line)[1])
		strength := calcStrength(cards, nil, false, 0)

		cardsByStrength[strength] = append(cardsByStrength[strength], Hand{cards, bid, 0})
	}

	for _, hands := range cardsByStrength {
		if len(hands) > 1 {
			sort.Slice(hands, func(i, j int) bool {
				return compareCards(hands[i].cards, hands[j].cards, getCardStrength)
			})
		}
	}

	var strengths []int
	for k := range cardsByStrength {
		strengths = append(strengths, k)
	}

	slices.Sort(strengths)

	var sum int
	rank := 1
	for _, s := range strengths {
		for _, hand := range cardsByStrength[s] {
			sum += hand.bid * rank
			rank++
		}
	}

	return sum
}

func partTwo(fileInput string) int {
	lines := strings.Split(fileInput, "\n")

	var cardsByStrength = make(map[int][]Hand)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		cards := strings.Fields(line)[0]
		bid, _ := strconv.Atoi(strings.Fields(line)[1])
		strength := calcStrength(cards, nil, true, 0)

		cardsByStrength[strength] = append(cardsByStrength[strength], Hand{cards, bid, 0})
	}

	for _, hands := range cardsByStrength {
		if len(hands) > 1 {
			sort.Slice(hands, func(i, j int) bool {
				return compareCards(hands[i].cards, hands[j].cards, getCardStrengthTwo)
			})
		}
	}

	var strengths []int
	for k := range cardsByStrength {
		strengths = append(strengths, k)
	}

	slices.Sort(strengths)

	var sum int
	rank := 1
	for _, s := range strengths {
		for _, hand := range cardsByStrength[s] {
			sum += hand.bid * rank
			rank++
		}
	}

	return sum
}

func compareCards(cardsOne, cardsTwo string, strengthFn func(c byte) int) bool {
	for i := 0; i < len(cardsOne); i++ {
		if strengthFn(cardsOne[i]) == strengthFn(cardsTwo[i]) {
			continue
		}

		return strengthFn(cardsOne[i]) < strengthFn(cardsTwo[i])
	}

	return false
}

func getCardStrength(card byte) int {
	return strings.Index("23456789TJQKA", string(card))
}

func getCardStrengthTwo(card byte) int {
	return strings.Index("J23456789TQKA", string(card))
}

func calcStrength(cards string, numOfPairs []int, withJokers bool, jokers int) int {
	if len(cards) == 0 {
		if withJokers && jokers > 0 {
			numOfPairs = replaceJokers(numOfPairs, jokers)
		}
		slices.Sort(numOfPairs)
		for i, rank := range ranks {
			if slices.Compare(rank, numOfPairs) == 0 {
				return i + 1
			}
		}
	}

	matches := 1
	compared := string(cards[0])
	if compared == "J" {
		jokers++
	}
	cards = cards[1:]

	for strings.Index(cards, compared) != -1 {
		matches++
		if compared == "J" {
			jokers++
		}
		cards = strings.Replace(cards, compared, "", 1)
	}
	numOfPairs = append(numOfPairs, matches)

	return calcStrength(cards, numOfPairs, withJokers, jokers)
}

func replaceJokers(numOfPairs []int, jokers int) []int {
	var jokersIndex = slices.Index(numOfPairs, jokers)
	if jokersIndex == -1 || jokers == 5 {
		return numOfPairs
	}
	numOfPairs = append(numOfPairs[:jokersIndex], numOfPairs[jokersIndex+1:]...)
	slices.Sort(numOfPairs)
	numOfPairs[len(numOfPairs)-1] += jokers

	return numOfPairs
}
