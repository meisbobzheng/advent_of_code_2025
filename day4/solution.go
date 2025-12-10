package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkAt(input string) int {
	if input == "@" {
		return 1
	} else {
		return 0
	}
}

func clearOut(graph [][]string, xLength int, yLength int) int {
	count := 0

	for y := range yLength {
		for x := range xLength {
			currentVal := graph[y][x]

			neighbors := 0

			if currentVal != "@" {
				continue
			}

			if x+1 < xLength {
				neighbors += checkAt(graph[y][x+1])

				if y+1 < yLength {
					neighbors += checkAt(graph[y+1][x+1])
				}

				if y-1 >= 0 {
					neighbors += checkAt(graph[y-1][x+1])
				}
			}

			if x-1 >= 0 {
				neighbors += checkAt(graph[y][x-1])

				if y+1 < yLength {
					neighbors += checkAt(graph[y+1][x-1])
				}

				if y-1 >= 0 {
					neighbors += checkAt(graph[y-1][x-1])
				}
			}

			if (y + 1) < yLength {
				neighbors += checkAt(graph[y+1][x])
			}

			if y-1 >= 0 {
				neighbors += checkAt(graph[y-1][x])
			}

			if neighbors < 4 {
				graph[y][x] = "X"
				count += 1
			}
		}
	}

	return count
}

func main() {
	var filepath string

	fmt.Print("Filepath: ")
	_, err := fmt.Scanln(&filepath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Received filepath: " + filepath)

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := [][]string{}
	xLength := 0
	yLength := 0
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		xLength = len(line)
		row := make([]string, len(line))

		for i, val := range line {
			row[i] = string(val)
		}

		graph = append(graph, row)
		yLength += 1
	}

	running := true

	for running {
		cleared := clearOut(graph, xLength, yLength)
		fmt.Printf("New loop output: %d\n", cleared)

		if cleared == 0 {
			running = false
		}

		count += cleared
	}

	fmt.Printf("Count: %d", count)

}
