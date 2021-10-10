package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	subscription, err := readFile("./data/match_patterns.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(subscription)
	networkSubscription, err := readNetworkFile("https://github.com/cobaltdisco/Google-Chinese-Results-Blocklist/raw/master/uBlacklist_match_patterns.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(networkSubscription)
	fmt.Println(removeDuplicate(subscription))
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

func removeDuplicate(arr []string) []string {
	start := time.Now()

	resArray := make([]string, 0)
	tmpArray := make(map[string]interface{})

	for _, val := range arr {
		if _, ok := tmpArray[val]; !ok {
			// 不存在则创建
			resArray = append(resArray, val)
			tmpArray[val] = struct{}{}
		}
	}
	fmt.Println("Removing duplicate time:", fmt.Sprintf("%vms", (time.Now().UnixNano()-start.UnixNano())/1e+6))
	return resArray
}
