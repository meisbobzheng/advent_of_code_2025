package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type FreshRange struct {
	Start int
	End   int
}

func main() {
	filePath := os.Args[1]
	fmt.Println("Received filepath", filePath)

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstHalf := true

	var freshRanges []FreshRange

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			firstHalf = false
			continue
		}

		if firstHalf {
			vals := strings.Split(line, "-")
			if len(vals) != 2 {
				continue
			}

			valInt1, err1 := strconv.Atoi(vals[0])
			valInt2, err2 := strconv.Atoi(vals[1])
			if err1 != nil || err2 != nil {
				log.Printf("Error parsing range: %s - %v, %v", line, err1, err2)
				continue
			}

			freshRanges = append(freshRanges, FreshRange{Start: valInt1, End: valInt2})

		} else {
			break
		}
	}

	slices.SortFunc(freshRanges, func(a, b FreshRange) int {
		if a.Start < b.Start {
			return -1
		} else if a.Start == b.Start {
			if a.End < b.End {
				return -1
			} else if a.End == b.End {
				return 0
			} else if a.End > b.End {
				return 1
			}
		} else if a.Start > b.Start {
			return 1
		}
		return 0
	})

	fmt.Println("Sorted ranges:", freshRanges)

	for i := 0; i < len(freshRanges)-1; i++ {
		if freshRanges[i].End >= freshRanges[i+1].Start {
			fmt.Println("Merging", freshRanges[i], freshRanges[i+1])
			freshRanges[i].End = max(freshRanges[i].End, freshRanges[i+1].End)
			freshRanges = slices.Delete(freshRanges, i+1, i+2)
			i--
		}
	}

	fmt.Println("Merged ranges:", freshRanges)

	sum := 0

	for i := 0; i < len(freshRanges); i++ {
		diff := freshRanges[i].End - freshRanges[i].Start + 1
		fmt.Println(freshRanges[i].Start, freshRanges[i].End, diff)

		sum += diff
	}

	fmt.Println(sum)
}

/**
p1 solution

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

	var freshRanges []FreshRange
	freshCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			firstHalf = false
			continue
		}

		if firstHalf {
			vals := strings.Split(line, "-")
			if len(vals) != 2 {
				continue
			}

			valInt1, err1 := strconv.Atoi(vals[0])
			valInt2, err2 := strconv.Atoi(vals[1])
			if err1 != nil || err2 != nil {
				log.Printf("Error parsing range: %s - %v, %v", line, err1, err2)
				continue
			}

			freshRanges = append(freshRanges, FreshRange{Start: valInt1, End: valInt2})

		} else {
			parsedVal, err := strconv.Atoi(line)
			if err != nil {
				log.Printf("Error parsing ingredient ID: %s - %v", line, err)
				continue
			}

			for _, r := range freshRanges {
				if parsedVal >= r.Start && parsedVal <= r.End {
					freshCount++
					break
				}
			}
		}
	}

	fmt.Println(freshCount)
}

**/
