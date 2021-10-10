package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	subscription, err := readFile("./data/match_patterns.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(subscription)
	networkSubscription, err := readNetworkFile("https://raw.githubusercontent.com/cobaltdisco/Google-Chinese-Results-Blocklist/master/uBlacklist_subscription.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(networkSubscription)
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

func readNetworkFile(path string) ([]string, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	fmt.Println("response: ", resp.Status)
	lines, err := ioToArr(resp.Body)
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
