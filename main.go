package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Finding struct {
	File    string
	Line    int
	Rule    string
	Preview string
}

func scanFile(path string) ([]Finding, error) {
	findings := []Finding{}
	lineNum := 1
	f, err := os.Open(path)
	if err != nil {
		return findings, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		currentLine := scanner.Text()
		header := strings.ToUpper(currentLine)
		if strings.Contains(header, "API_KEY") {
			finding := Finding{
				File:    path,
				Line:    lineNum,
				Rule:    "API_KEY",
				Preview: "",
			}
			findings = append(findings, finding)
		}
		lineNum++
	}
	err = scanner.Err()
	if err != nil {
		return findings, err
	}
	return findings, nil
}

func maskValue(value string) string {
	if len(value) <= 12 {
		return "****"
	} else {
		return value[0:4] + "****"
	}
}

func extractValue(line string) (string, string, bool) {
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "#")
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "export ")
	key, value, found := strings.Cut(line, "=")
	if !found {
		return "", "", false
	}
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)
	value = strings.Trim(value, `"'`)
	if key == "" || value == "" {
		return "", "", false
	}
	return key, value, true
}

func main() {
	findings, err := scanFile("testdata/sample.env")
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of findings:", len(findings))
	test := []string{"API_KEY=abc123", "# API_KEY=abc123", "# export API_KEY=abc123", `export API_KEY="abc123"`, "API_KEY='abc123'", "  API_KEY = abc123  ", "DB_URL=postgres://user:pw@host/db?sslmode=require", "API_KEY=", `API_KEY=""`, "just some text", "=abc"}
	for _, line := range test {
		key, value, found := extractValue(line)
		fmt.Printf("input=%q  key=%q  value=%q  found=%v\n", line, key, value, found)
		if found {
			fmt.Printf("  preview=%q\n", maskValue(value))
		}
	}
}
