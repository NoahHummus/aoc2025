package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	IDRanges := ParseInput()

	var invalidIDList []int
	for counter, IDRange := range IDRanges {
		log.Print(counter)
		elements := strings.Split(IDRange, "-")
		startRange, _ := strconv.Atoi(elements[0])
		endRange, _ := strconv.Atoi(elements[1])

		for id := startRange; id <= endRange; id++ {
			log.Print("Proccessing id -", id)
			digitCounterMap := make(map[int]int)
			var invalidID bool
			idString := strconv.Itoa(id)
			for _, digit := range idString {
				digitInt, _ := strconv.Atoi(digit)

				_, exists := digitCounterMap[digit]
				if exists {
					digitCounterMap[digit] += 1
				} else {
					digitCounterMap[digit] = 1
				}

			}

			for _, value := range digitCounterMap {
				if value != 2 {
					invalidID = false

				}

			}

			if invalidID {
				invalidIDList = append(invalidIDList, id)

			}

		}

	}

	log.Print(invalidIDList)

}

func ParseInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elements := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		elements = strings.Split(line, ",")

	}
	return elements

}
