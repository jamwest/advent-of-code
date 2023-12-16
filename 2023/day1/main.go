package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partOneLine(line []byte) int {
	sum := 0

	for i := 0; i < len(line); i++ {
		if line[i] >= 48 && line[i] <= 58 {
			sum = int(line[i]-48) * 10
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= 48 && line[i] <= 58 {
			sum = sum + int(line[i]-48)
			break
		}
	}
	return sum
}

func partOne(scanner *bufio.Scanner) int {
	sum := 0

	for scanner.Scan() {
		sum = sum + partOneLine(scanner.Bytes())
	}
	return sum
}

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	buf := bufio.NewScanner(f)
	partOne := partOne(buf)
	check(buf.Err())
	println(partOne)
}
