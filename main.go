package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	subscription, err := readFile("./data/match_patterns.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(subscription)
}

// 读取文件到字符串数组
func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
