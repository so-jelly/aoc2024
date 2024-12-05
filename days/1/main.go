package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/so-jelly/aoc2024/pkg/maths"
	"github.com/so-jelly/aoc2024/pkg/read"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	part1, err := part1(read.ScanString(string(input)))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", part1)
	part2, err := part2(read.ScanString(string(input)))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2: %d\n", part2)
}

func part1(scanner *bufio.Scanner) (distance int64, err error) {
	// columns just need to be sorted
	var lefts, rights []int64
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Fields(line)
		left, err := strconv.ParseInt(pair[0], 10, 64)
		if err != nil {
			return 0, err
		}
		right, err := strconv.ParseInt(pair[1], 10, 64)
		if err != nil {
			return 0, err
		}
		lefts = append(lefts, left)
		rights = append(rights, right)
	}
	slices.Sort(lefts)
	slices.Sort(rights)
	if len(lefts) != len(rights) {
		return 0, nil
	}
	for i := 0; i < len(lefts); i++ {
		distance += maths.Abs64(rights[i] - lefts[i])
	}
	return distance, nil
}

func part2(scanner *bufio.Scanner) (distance int64, err error) {

	var lefts []int64
	var rights = make(map[int64]int64)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Fields(line)
		left, err := strconv.ParseInt(pair[0], 10, 64)
		if err != nil {
			return 0, err
		}
		lefts = append(lefts, left)

		right, err := strconv.ParseInt(pair[1], 10, 64)
		if err != nil {
			return 0, err
		}
		rights[right]++
	}
	var score int64
	for i := 0; i < len(lefts); i++ {
		if occurrences, ok := rights[lefts[i]]; ok {
			score += lefts[i] * occurrences
		}
	}
	return score, nil
}
