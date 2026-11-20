package main

import (
	"bufio"
	"fmt"
	"os"
)

type Finding struct {
	File    string
	Line    int
	Rule    string
	Preview string
}

func scanFile(path string) ([]Finding, error) {
	findings := []Finding{}
	counter := 1
	f, err := os.Open(path)
	if err != nil {
		return findings, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counter++
	}
	err = scanner.Err()
	if err != nil {
		return findings, err
	}
	return findings, nil
}

func main() {
	find, err := scanFile("testdata/sample.env")
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of findings:", len(find))
}
