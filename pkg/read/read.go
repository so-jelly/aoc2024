package read

import (
	"bufio"
	"os"
	"strings"
)

func ScanStdin() *bufio.Scanner {
	return bufio.NewScanner(bufio.NewReader(os.Stdin))
}

func ScanFile(file string) (*bufio.Scanner, error) {
	f, err := os.Open(file)
	return bufio.NewScanner(f), err
}

func ScanString(input string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(input))
}
