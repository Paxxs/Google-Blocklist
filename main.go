package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	subscription, err := readFile("./data/match_patterns.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(subscription)
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	lines, err := ioToArr(file)
	return lines, err
}

func ioToArr(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
