package main

import (
	"fmt"
	//"sort"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := "input.txt"

	data, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	ranges := strings.Split(string(data), ",")
	sum := 0

	for _, val := range ranges {
		splitRange := strings.Split(strings.TrimSpace(val), "-")

		if len(splitRange) != 2 {
			log.Fatal("Invalid range")
		}

		num1, err := strconv.Atoi(splitRange[0])
		if err != nil {
			log.Fatal(err)
		}

		num2, err := strconv.Atoi(splitRange[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := num1; i <= num2; i++ {
			numStr := strconv.Itoa(i)

			repeatedStr := strings.Repeat(numStr, 2)

			if strings.Contains(repeatedStr[1:len(repeatedStr)-1], numStr) {
				fmt.Printf("Found: %d\n", i)
				sum += i
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum)

}
