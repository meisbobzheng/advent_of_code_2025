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

	filepath := "input.txt"
	// filepath := "sample.txt"

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Failed to open file: %s\n", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Original line:%s\n", line)

		var solution [12]string
		lastIndex := 0
		padding := len(line) - 12

		for i := range 12 {
			largest := -1
			largestIndex := lastIndex

			// fmt.Printf("Padding: %d\n", padding)
			// fmt.Printf("Last index: %d\n", lastIndex)

			if padding == 0 {
				charVal := line[lastIndex]

				intVal, err := strconv.Atoi(string(charVal))

				if err != nil {
					log.Fatal("Failed to convert char to int")
				}

				largest = intVal

			} else {
				for v := lastIndex; v <= lastIndex+padding; v++ {
					charVal := line[v]
					intVal, err := strconv.Atoi(string(charVal))

					if err != nil {
						log.Fatal("Failed to convert char to int")
					}

					if intVal > largest {
						largestIndex = v
						largest = intVal
					}
				}

			}

			solution[i] = strconv.Itoa(largest)
			padding = padding - (largestIndex - lastIndex)
			lastIndex = largestIndex + 1
		}

		longestLine := strings.Join(solution[:], "")

		parsedLine, err := strconv.Atoi(longestLine)

		if err != nil {
			log.Fatalf("Failed to parse line: %s\n", err)
		}

		sum += parsedLine

		fmt.Printf("Longest 12: %s\n", longestLine)

	}

	fmt.Printf("Sum: %d\n", sum)
}

// p1 solution
// largest := 0
// 		second := -1
//
// 		for index, char := range line {
// 			charValue, err := strconv.Atoi(string(char))
//
// 			if err != nil {
// 				log.Fatal("Failed to convert string to int")
// 			}
//
// 			if (index != len(line)-1) && charValue > largest {
// 				largest = charValue
// 				second = -1
// 				continue
// 			}
//
// 			if charValue > second {
// 				second = charValue
// 			}
//
// 		}
//
// 		sum += (10 * largest) + second
//
