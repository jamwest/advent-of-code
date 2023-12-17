package main

import (
	"bufio"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partOne(line []byte) int {
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

func convertDigitString(digitString string) int {
	var result int
	switch digitString {
	case "one":
		result = 1
	case "two":
		result = 2
	case "three":
		result = 3
	case "four":
		result = 4
	case "five":
		result = 5
	case "six":
		result = 6
	case "seven":
		result = 7
	case "eight":
		result = 8
	case "nine":
		result = 9
	default:
		panic(nil)
	}
	return result
}

func reverseByteArray(s []byte) []byte {
	output := s
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	return output
}

func partTwo(line []byte) int {
	sum := 0

	re := regexp.MustCompile("1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine")
	reReverse := regexp.MustCompile("enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|9|8|7|6|5|4|3|2|1")
	index := re.FindIndex(line)

	if index == nil {
		panic(nil)
	}
	first := line[index[0]:index[1]]

	if len(first) == 1 {
		sum = int(first[0]-48) * 10
	} else {
		sum = convertDigitString(string(first)) * 10
	}

	lineReverse := reverseByteArray(line)
	index = reReverse.FindIndex(lineReverse)

	if index == nil {
		panic(nil)
	}

	last := lineReverse[index[0]:index[1]]

	if len(last) == 1 {
		sum = sum + int(last[0]-48)
	} else {
		sum = sum + convertDigitString(string(reverseByteArray(last)))
	}

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
