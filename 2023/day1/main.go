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

func main() {
	f, err := os.Open("./input.txt")
	check(err)

	buf := bufio.NewReader(f)

	line, _, err := buf.ReadLine()
	check(err)

	println(string(line))

	f.Close()
}
