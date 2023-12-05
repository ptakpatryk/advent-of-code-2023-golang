package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/ptakpatryk/advent-of-code-2023-golang/lib"
)

var input, err = os.ReadFile("day_05/input.txt")

type NumberMap struct {
	DestRangeStart   int
	SourceRangeStart int
	RangeLenght      int
}

func main() {
	lib.CheckError(err)
	fmt.Println(partOne(string(input)))
	fmt.Println(partTwo(string(input)))
}

func partOne(fileInput string) int {
	split := strings.Split(fileInput, "\n\n")
	strings.Split(split[0], "")
	seeds := strings.Fields(strings.Split(split[0], ":")[1])
	mapPipe := parseMaps(split[1:])

	var minLocation int

	for _, seed := range seeds {
		location, _ := strconv.Atoi(seed)

		locationAfterMapping := runThroughMaps(location, mapPipe)
		if minLocation == 0 || locationAfterMapping < minLocation {
			minLocation = locationAfterMapping
		}
	}

	return minLocation
}

func partTwo(fileInput string) int {
	split := strings.Split(fileInput, "\n\n")
	strings.Split(split[0], "")
	seedsRanges := strings.Fields(strings.Split(split[0], ":")[1])
	mapPipe := parseMaps(split[1:])

	var wg sync.WaitGroup
	var mu sync.Mutex
	var minLocation int

	for i := 0; i < len(seedsRanges); i += 2 {
		wg.Add(1)
		start, _ := strconv.Atoi(seedsRanges[i])
		length, _ := strconv.Atoi(seedsRanges[i+1])

		go func() {
			defer wg.Done()
			var rangeMinLocation int

			for i := 0; i < length; i++ {
				seed := start + i
				location := seed

				locationAfterMapping := runThroughMaps(location, mapPipe)

				if rangeMinLocation == 0 || locationAfterMapping < rangeMinLocation {
					rangeMinLocation = locationAfterMapping
				}
			}

			mu.Lock()
			defer mu.Unlock()

			if minLocation == 0 || rangeMinLocation < minLocation {
				minLocation = rangeMinLocation
			}
		}()
	}

	wg.Wait()

	return minLocation
}

func runThroughMaps(seed int, pipe map[int][]NumberMap) int {
	location := seed
	for i := 0; i <= len(pipe); i++ {
		for _, mapVal := range pipe[i] {
			if location >= mapVal.SourceRangeStart && location < mapVal.SourceRangeStart+mapVal.RangeLenght {
				diff := location - mapVal.SourceRangeStart
				location = mapVal.DestRangeStart + diff
				break
			}
		}
	}

	return location
}

func parseMaps(maps []string) map[int][]NumberMap {
	var mapPipe = make(map[int][]NumberMap)

	for i, mapValues := range maps {
		values := strings.Split(mapValues, "\n")[1:]
		for _, mapVal := range values {
			if len(mapVal) == 0 {
				continue
			}
			input := strings.Fields(mapVal)
			drs, _ := strconv.Atoi(input[0])
			srs, _ := strconv.Atoi(input[1])
			rl, _ := strconv.Atoi(input[2])
			mapPipe[i] = append(mapPipe[i], NumberMap{drs, srs, rl})
		}
	}

	return mapPipe
}
