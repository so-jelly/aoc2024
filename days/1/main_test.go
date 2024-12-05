package main

import (
	"os"
	"testing"

	"github.com/so-jelly/aoc2024/pkg/read"
)

var testInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func Test_part1(t *testing.T) {
	distance := int64(11)
	got, err := part1(read.ScanString(testInput))
	if err != nil {
		t.Errorf("part1() = %v, want %v", got, distance)
	}
	if got != distance {
		t.Errorf("part1() = %v, want %v", got, distance)
	}

	distance = int64(3574690)
	scanner, err := read.ScanFile("testdata.txt")
	if err != nil {
		t.Errorf("failed to open file: %v", err)
	}

	got, err = part1(scanner)
	if err != nil {
		t.Errorf("part1() = %v, want %v", got, distance)
	}
	if got != distance {
		t.Errorf("part1() = %v, want %v", got, distance)
	}

}

func Test_part2(t *testing.T) {

	score := int64(31)
	got, err := part2(read.ScanString(testInput))
	if err != nil {
		t.Errorf("part2() = %v, want %v", got, score)
	}
	if got != score {
		t.Errorf("part2() = %v, want %v", got, score)
	}

	score = int64(22565391)
	scanner, err := read.ScanFile("testdata.txt")
	if err != nil {
		t.Errorf("failed to open file: %v", err)
	}
	got, err = part2(scanner)
	if err != nil {
		t.Errorf("part2() = %v, want %v", got, score)
	}
	if got != score {
		t.Errorf("part2() = %v, want %v", got, score)
	}
}

func Benchmark_part1(b *testing.B) {
	d, err := os.ReadFile("testdata.txt")
	if err != nil {
		b.Errorf("failed to open file: %v", err)
	}
	for i := 0; i < b.N; i++ {
		part1(read.ScanString(string(d)))
	}
}
