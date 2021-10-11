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
	sub_original, err := loadSubscriptions("./config.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println("original subscriptions size: ", len(sub_original))
	sub := removeDuplicate(sub_original)
	fmt.Println("removeDuplicated size: ", len(sub))
	err = writeFile("uBlacklist_subscription.txt", sub)
	if err != nil {
		panic(err)
	}
	fmt.Println("done!")
}

func writeFile(path string, contents []string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	datewriter := bufio.NewWriter(file)
	for _, lines := range contents {
		datewriter.WriteString(lines + "\n")
	}
	datewriter.Flush()

	return file.Close()
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
