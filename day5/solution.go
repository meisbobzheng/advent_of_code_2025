package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := os.Args[1]
	fmt.Println("Received filepath", filePath)

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	firstHalf := true

	ranges := make(map[int]int)
	freshCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			firstHalf = false
			continue
		}

		if firstHalf {
			vals := strings.Split(line, "-")
			valInt1, _ := strconv.Atoi(vals[0])
			valInt2, _ := strconv.Atoi(vals[1])
			ranges[valInt1] = valInt2
		} else {
			parsedVal, _ := strconv.Atoi(line)

			for start, end := range ranges {
				if parsedVal >= start && parsedVal <= end {
					fmt.Println(start, end, parsedVal)
					freshCount++
					break
				}
			}

		}
	}

	fmt.Println(freshCount)
}
