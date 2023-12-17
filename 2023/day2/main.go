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

func partOne(line []byte) int {
	sum := 0
	return sum
}

func partTwo(line []byte) int {
	sum := 0

	return sum
}

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	buf := bufio.NewScanner(f)
	sum1 := 0
	sum2 := 0
	for buf.Scan() {
		line := buf.Bytes()
		sum1 = sum1 + partOne(line)
		sum2 = sum2 + partTwo(line)
	}

	check(buf.Err())

	println(sum1)
	println(sum2)
}
